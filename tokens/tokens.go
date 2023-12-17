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
	"github.com/unpackdev/solgo/exchanges"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

type Token struct {
	ctx             context.Context
	network         utils.Network
	networkID       utils.NetworkID
	bindManager     *bindings.Manager
	exchangeManager *exchanges.Manager
	clientPool      *clients.ClientPool
	simulator       *simulator.Simulator
	tokenBind       *bindings.Token
	simulatorType   utils.SimulatorType
	inSimulation    atomic.Bool
	descriptor      *Descriptor
	exchanges       map[utils.ExchangeType]map[utils.SimulatorType]exchanges.Exchange
}

func NewToken(ctx context.Context, network utils.Network, address common.Address, simulatorType utils.SimulatorType, bindManager *bindings.Manager, exchangeManager *exchanges.Manager, sim *simulator.Simulator, clientPool *clients.ClientPool) (*Token, error) {
	if bindManager == nil {
		return nil, errors.New("bind manager is nil")
	}

	if clientPool == nil {
		return nil, errors.New("client pool is nil")
	}

	toReturn := &Token{
		ctx:             ctx,
		network:         network,
		networkID:       utils.GetNetworkID(network),
		bindManager:     bindManager,
		exchangeManager: exchangeManager,
		clientPool:      clientPool,
		simulator:       sim,
		inSimulation:    atomic.Bool{},
		simulatorType:   simulatorType,
		descriptor: &Descriptor{
			Address: address,
			Pairs:   make(map[utils.ExchangeType]*Pair),
		},
		exchanges: make(map[utils.ExchangeType]map[utils.SimulatorType]exchanges.Exchange),
	}

	if err := toReturn.PrepareBindings(ctx); err != nil {
		return nil, fmt.Errorf("failed to prepare bindings: %w", err)
	}

	return toReturn, nil
}

func (t *Token) GetDescriptor() *Descriptor {
	return t.descriptor
}

func (t *Token) GetContext() context.Context {
	return t.ctx
}

func (t *Token) GetNetwork() utils.Network {
	return t.network
}

func (t *Token) GetNetworkID() utils.NetworkID {
	return t.networkID
}

func (t *Token) GetBindManager() *bindings.Manager {
	return t.bindManager
}

func (t *Token) GetClientPool() *clients.ClientPool {
	return t.clientPool
}

func (t *Token) GetSimulator() *simulator.Simulator {
	return t.simulator
}

func (t *Token) GetSimulatorType() utils.SimulatorType {
	return t.simulatorType
}

func (t *Token) GetExchangeManager() *exchanges.Manager {
	return t.exchangeManager
}

func (t *Token) IsInSimulation() bool {
	return t.inSimulation.Load()
}

func (t *Token) SetInSimulation(inSimulation bool) {
	t.inSimulation.Store(inSimulation)
}

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

func (t *Token) GetClient(ctx context.Context) (*clients.Client, error) {
	var client *clients.Client
	var err error

	if t.IsInSimulation() {
		client, _, err = t.simulator.GetClient(ctx, t.simulatorType, t.descriptor.BlockNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to get client from simulator: %s", err)
		}
	} else {
		client = t.clientPool.GetClientByGroup(t.network.String())
		if client == nil {
			return nil, fmt.Errorf("failed to get client from client pool: %s", err)
		}
	}

	return client, err
}

func (t *Token) GetSimulatedClient(ctx context.Context, simulatorType utils.SimulatorType, atBlock *big.Int) (*clients.Client, error) {
	if t.simulator == nil {
		return nil, errors.New("simulator is not set")
	}

	block := t.descriptor.BlockNumber

	if atBlock != nil {
		block = atBlock
	}

	client, _, err := t.simulator.GetClient(ctx, simulatorType, block)
	if err != nil {
		return nil, fmt.Errorf("failed to get client from simulator: %s - block number: %v", err, block)
	}

	return client, err
}

func (t *Token) Unpack(ctx context.Context, atBlock *big.Int, simulate bool, simulatorType utils.SimulatorType) (*Descriptor, error) {
	var tokenBinding *bindings.Token
	var err error

	t.descriptor.BlockNumber = atBlock

	if simulate {
		t.SetInSimulation(true)
		simBind, _ := t.simulator.GetBindingManager(simulatorType)
		tokenBinding, err = t.GetTokenBind(ctx, simulatorType, simBind, atBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to get simulated token bindings: %w", err)
		}
	} else {
		tokenBinding, err = t.GetTokenBind(ctx, simulatorType, t.bindManager, atBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to get token bindings: %w", err)
		}
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
