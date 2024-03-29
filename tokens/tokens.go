package tokens

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

// Token encapsulates the necessary details and operations for interacting with an Ethereum token.
type Token struct {
	ctx           context.Context     // The context for network requests.
	network       utils.Network       // The blockchain network the token exists on.
	networkID     utils.NetworkID     // The numeric ID of the network.
	bindManager   *bindings.Manager   // The contract binding manager.
	clientPool    *clients.ClientPool // The pool of Ethereum network clients.
	tokenBind     *bindings.Token     // The smart contract bindings for the token.
	simulatorType utils.SimulatorType // The type of simulator used for testing.
	inSimulation  atomic.Bool         // Flag indicating if the token is being simulated.
	descriptor    *Descriptor         // The token's descriptor containing detailed information.
}

// NewToken creates a new Token instance with the specified parameters.
// It prepares the contract bindings and validates the provided inputs.
func NewToken(ctx context.Context, network utils.Network, address common.Address, bindManager *bindings.Manager, pool *clients.ClientPool) (*Token, error) {
	if bindManager == nil {
		return nil, errors.New("bind manager is nil")
	}

	if pool == nil {
		return nil, errors.New("client pool is nil")
	}

	toReturn := &Token{
		ctx:          ctx,
		network:      network,
		networkID:    utils.GetNetworkID(network),
		bindManager:  bindManager,
		clientPool:   pool,
		inSimulation: atomic.Bool{},
		descriptor: &Descriptor{
			Address: address,
		},
	}

	if err := toReturn.PrepareBindings(ctx); err != nil {
		return nil, fmt.Errorf("failed to prepare bindings: %w", err)
	}

	return toReturn, nil
}

// GetDescriptor returns the Descriptor of the token containing its metadata.
func (t *Token) GetDescriptor() *Descriptor {
	return t.descriptor
}

// GetContext returns the context associated with the Token instance.
func (t *Token) GetContext() context.Context {
	return t.ctx
}

// GetNetwork returns the blockchain network the token is associated with.
func (t *Token) GetNetwork() utils.Network {
	return t.network
}

// GetNetworkID returns the numeric ID of the blockchain network.
func (t *Token) GetNetworkID() utils.NetworkID {
	return t.networkID
}

// GetBindManager returns the contract binding manager associated with the Token.
func (t *Token) GetBindManager() *bindings.Manager {
	return t.bindManager
}

// GetClientPool returns the pool of Ethereum network clients.
func (t *Token) GetClientPool() *clients.ClientPool {
	return t.clientPool
}

// GetSimulatorType returns the type of simulator used for testing.
func (t *Token) GetSimulatorType() utils.SimulatorType {
	return t.simulatorType
}

// IsInSimulation returns true if the token is currently being simulated.
func (t *Token) IsInSimulation() bool {
	return t.inSimulation.Load()
}

// GetBind returns the smart contract bindings for the token.
func (t *Token) GetBind() *bindings.Token {
	return t.tokenBind
}

// SetInSimulation sets the simulation flag for the token.
func (t *Token) SetInSimulation(inSimulation bool) {
	t.inSimulation.Store(inSimulation)
}

// GetEntity returns the token entity, constructing it if necessary.
func (t *Token) GetEntity() *entities.Token {
	if t.descriptor.Entity == nil {
		t.descriptor.Entity = entities.NewToken(
			uint(utils.GetNetworkID(t.network)),
			t.descriptor.Address,
			uint(t.descriptor.Decimals),
			t.descriptor.Symbol,
			t.descriptor.Name,
		)
	}

	return t.descriptor.Entity
}

// GetClient retrieves an Ethereum client from the client pool.
func (t *Token) GetClient() (*clients.Client, error) {
	client := t.clientPool.GetClientByGroup(t.network.String())
	if client == nil {
		return nil, errors.New("failed to get client from client pool")
	}
	return client, nil
}

// Unpack resolves the token's details (name, symbol, decimals, total supply) at a specific block number.
func (t *Token) Unpack(ctx context.Context, atBlock *big.Int) (*Descriptor, error) {
	var tokenBinding *bindings.Token
	var err error

	t.descriptor.BlockNumber = atBlock

	tokenBinding, err = t.GetTokenBind(ctx, t.bindManager)
	if err != nil {
		return nil, fmt.Errorf("failed to get token bindings: %w", err)
	}
	t.tokenBind = tokenBinding

	t.descriptor.Name, err = t.ResolveName(ctx, t.descriptor.Address, tokenBinding)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token name: %w", err)
	}

	t.descriptor.Symbol, err = t.ResolveSymbol(ctx, t.descriptor.Address, tokenBinding)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token symbol: %w", err)
	}

	t.descriptor.Decimals, err = t.ResolveDecimals(ctx, t.descriptor.Address, tokenBinding)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token decimals: %w", err)
	}

	t.descriptor.TotalSupply, err = t.ResolveTotalSupply(ctx, t.descriptor.Address, tokenBinding)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token total supply: %w", err)
	}

	t.descriptor.Entity = t.GetEntity()

	return t.descriptor, nil
}
