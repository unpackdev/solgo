package simulator

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

// Simulator is the core struct for the simulator package. It manages the lifecycle
// and operations of blockchain simulations, including the management of providers
// and faucets. The Simulator allows for flexible interaction with various simulated
// blockchain environments and accounts.
type Simulator struct {
	ctx        context.Context                  // Context for managing lifecycle and control flow.
	opts       *Options                         // Configuration options for the Simulator.
	clientPool *clients.ClientPool              // Ethereum client pool
	providers  map[utils.SimulatorType]Provider // Registered providers for different simulation types.
	mu         sync.Mutex                       // Mutex to protect concurrent access to the providers map and other shared resources.
	faucets    *Faucet                          // Faucet for managing simulated accounts.
}

// NewSimulator initializes a new Simulator with the given context and options.
// It sets up a new Faucet for account management and prepares the Simulator for operation.
// Returns an error if options are not provided or if the Faucet fails to initialize.
func NewSimulator(ctx context.Context, opts *Options) (*Simulator, error) {
	if opts == nil {
		return nil, fmt.Errorf("in order to create a new simulator, you must provide options")
	}

	pool, err := clients.NewClientPool(ctx, &clients.Options{Nodes: []clients.Node{}})
	if err != nil {
		return nil, fmt.Errorf("failed to create simulator client pool: %w", err)
	}

	faucets, err := NewFaucet(ctx, pool, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create new faucet: %w", err)
	}

	if opts.FaucetsEnabled {
		if opts.FaucetAccountCount <= 0 {
			return nil, fmt.Errorf("auto faucet count must be greater than 0")
		}

		for _, network := range opts.SupportedNetworks {
			if faucetCount := len(faucets.List(network)); faucetCount < opts.FaucetAccountCount {
				missingFaucetCount := opts.FaucetAccountCount - faucetCount

				zap.L().Info(
					"Generating new faucet accounts. Please be patient...",
					zap.Int("current_count", faucetCount),
					zap.Int("required_count", opts.FaucetAccountCount),
					zap.Int("missing_count", missingFaucetCount),
					zap.String("network", string(network)),
				)

				for i := 0; i < missingFaucetCount; i++ {
					zap.L().Debug(
						"Generating new faucet account...",
						zap.Int("current_count", i),
						zap.Int("required_count", opts.FaucetAccountCount),
						zap.String("network", string(network)),
					)

					if _, err := faucets.Create(network, opts.DefaultPassword, true, utils.SimulatorAccountType.String()); err != nil {
						return nil, fmt.Errorf("failed to generate faucet account for network: %s err: %s", network, err)
					}
				}
			}
		}
	}

	return &Simulator{
		ctx:        ctx,
		opts:       opts,
		providers:  make(map[utils.SimulatorType]Provider),
		faucets:    faucets,
		clientPool: pool,
	}, nil
}

// GetFaucet returns the Faucet associated with the Simulator.
// The Faucet is used for creating and managing simulated blockchain accounts.
func (s *Simulator) GetFaucet() *Faucet {
	return s.faucets
}

// GetClientPool returns the Ethereum client pool associated with the Simulator.
// The client pool is used for managing Ethereum clients for various simulated environments.
func (s *Simulator) GetClientPool() *clients.ClientPool {
	return s.clientPool
}

func (c *Simulator) GetBindingManager(simType utils.SimulatorType) (*bindings.Manager, error) {
	return bindings.NewManager(c.ctx, c.GetProvider(simType).GetClientPool())
}

// RegisterProvider registers a new provider for a specific simulation type.
// If a provider for the given name already exists, it returns false.
// Otherwise, it adds the provider and returns true.
func (s *Simulator) RegisterProvider(name utils.SimulatorType, provider Provider) (bool, error) {
	if _, ok := s.providers[name]; ok {
		return false, fmt.Errorf("provider %s already exists", name)
	}

	s.providers[name] = provider

	if anvilProvider, ok := ToAnvilProvider(provider); ok {
		manager, err := s.GetBindingManager(name)
		if err != nil {
			return false, err
		}

		anvilProvider.SetBindingManager(manager)
		return true, nil
	}

	return true, nil
}

// GetProvider retrieves a registered provider by its simulation type.
// Returns nil if no provider is registered under the given name.
func (s *Simulator) GetProvider(name utils.SimulatorType) Provider {
	s.mu.Lock()
	defer s.mu.Unlock()

	if provider, ok := s.providers[name]; ok {
		return provider
	}

	return nil
}

// GetProviders returns a map of all registered providers by their simulation types.
func (s *Simulator) GetProviders() map[utils.SimulatorType]Provider {
	return s.providers
}

// ProviderExists checks if a provider with the given simulation type is registered.
// Returns true if the provider exists, false otherwise.
func (s *Simulator) ProviderExists(name utils.SimulatorType) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.providers[name]; ok {
		return true
	}

	return false
}

// Start initiates the simulation providers within the Simulator. It can start
// all registered providers or a subset specified in the 'simulators' variadic argument.
// If the 'simulators' argument is provided, only the providers matching the specified
// SimulatorTypes are started. If no 'simulators' argument is provided, all registered
// providers are started.
//
// This method iterates through each registered provider and calls its Load method,
// passing the provided context. The Load method of each provider is expected to
// initiate any necessary operations to start the simulation client.
func (s *Simulator) Start(ctx context.Context, simulators ...utils.SimulatorType) error {
	for _, provider := range s.providers {
		if len(simulators) > 0 {
			for _, simulator := range simulators {
				if provider.Type() == simulator {
					if err := provider.Load(ctx); err != nil {
						return fmt.Errorf("failed to start provider: %w", err)
					}
				}
			}
		} else {
			if err := provider.Load(ctx); err != nil {
				return fmt.Errorf("failed to start provider: %w", err)
			}
		}
	}

	return nil
}

// Stop terminates the simulation providers within the Simulator. Similar to Start,
// it can stop all registered providers or a subset specified in the 'simulators'
// variadic argument. If the 'simulators' argument is provided, only the providers
// matching the specified SimulatorTypes are stopped. If no 'simulators' argument
// is provided, all registered providers are stopped.
//
// This method iterates through each registered provider and calls its Unload method,
// passing the provided context. The Unload method of each provider is expected to
// perform any necessary operations to stop the simulation client.
func (s *Simulator) Stop(ctx context.Context, simulators ...utils.SimulatorType) error {
	for _, provider := range s.providers {
		if len(simulators) > 0 {
			for _, simulator := range simulators {
				if provider.Type() == simulator {
					if err := provider.Unload(ctx); err != nil {
						return err
					}
				}
			}
		} else {
			if err := provider.Unload(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

// Status retrieves the status of the simulation providers within the Simulator.
// Similar to Start and Stop, it can retrieve the status of all registered providers
// or a subset specified in the 'simulators' variadic argument. If the 'simulators'
// argument is provided, only the providers matching the specified SimulatorTypes
// are queried. If no 'simulators' argument is provided, all registered providers
// are queried.
//
// This method iterates through each registered provider and calls its Status method,
// passing the provided context. The Status method of each provider is expected to
// return the current status of the simulation client.
func (s *Simulator) Status(ctx context.Context, simulators ...utils.SimulatorType) (map[utils.SimulatorType][]*NodeStatus, error) {
	statuses := make(map[utils.SimulatorType][]*NodeStatus)

	for _, provider := range s.providers {
		if len(simulators) > 0 {
			for _, simulator := range simulators {
				if provider.Type() == simulator {
					statusSlice, err := provider.Status(ctx)
					if err != nil {
						return nil, err
					}
					for _, status := range statusSlice {
						statuses[provider.Type()] = append(statuses[provider.Type()], status)
					}
				}
			}
		} else {
			statusSlice, err := provider.Status(ctx)
			if err != nil {
				return nil, err
			}
			for _, status := range statusSlice {
				statuses[provider.Type()] = append(statuses[provider.Type()], status)
			}
		}
	}

	return statuses, nil
}

// GetClient retrieves a blockchain client for a given provider and block number.
func (s *Simulator) GetClient(ctx context.Context, provider utils.SimulatorType, blockNumber *big.Int) (*clients.Client, error) {
	if !s.ProviderExists(provider) {
		return nil, fmt.Errorf("provider %s does not exist", provider)
	}

	if providerCtx, ok := s.GetProvider(provider).(*AnvilProvider); ok {
		if node, ok := providerCtx.GetNodeByBlockNumber(blockNumber); !ok {
			// TODO: Spawn the new node or attempt to fetch the node....
			return nil, fmt.Errorf("node for block number %d does not exist", blockNumber)
		} else {
			if client, found := providerCtx.GetClientByGroupAndType(node.GetProvider().Type(), node.GetID().String()); found {
				return client, nil
			} else {
				return nil, fmt.Errorf("client for provider: %s - node %s does not exist", node.Provider, node.GetID().String())
			}
		}
	}

	return nil, fmt.Errorf("provider %s is not fully implemented", provider)
}
