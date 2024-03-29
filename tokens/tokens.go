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

type Token struct {
	ctx           context.Context
	network       utils.Network
	networkID     utils.NetworkID
	bindManager   *bindings.Manager
	clientPool    *clients.ClientPool
	tokenBind     *bindings.Token
	simulatorType utils.SimulatorType
	inSimulation  atomic.Bool
	descriptor    *Descriptor
}

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

func (t *Token) GetSimulatorType() utils.SimulatorType {
	return t.simulatorType
}

func (t *Token) IsInSimulation() bool {
	return t.inSimulation.Load()
}

func (t *Token) GetBind() *bindings.Token {
	return t.tokenBind
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

func (t *Token) GetClient() (*clients.Client, error) {
	client := t.clientPool.GetClientByGroup(t.network.String())
	if client == nil {
		return nil, errors.New("failed to get client from client pool")
	}
	return client, nil
}

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
