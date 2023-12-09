package exchanges

import (
	"context"
	"fmt"

	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/utils"
)

type Manager struct {
	ctx        context.Context
	clientPool *clients.ClientPool
	opts       *Options
	exchanges  map[utils.ExchangeType]Exchange
}

func NewManager(ctx context.Context, clientPool *clients.ClientPool, opts *Options) (*Manager, error) {
	if clientPool == nil {
		return nil, fmt.Errorf("client pool cannot be nil")
	}

	if opts == nil {
		return nil, fmt.Errorf("options cannot be nil")
	}

	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Manager{
		ctx:        ctx,
		clientPool: clientPool,
		opts:       opts,
		exchanges:  make(map[utils.ExchangeType]Exchange),
	}, nil
}

// RegisterExchange registers a new exchange.
func (m *Manager) RegisterExchange(name utils.ExchangeType, bindManager *bindings.Manager, exchangeFn exchangeFn) error {
	if _, ok := m.exchanges[name]; ok {
		return fmt.Errorf("exchange %s already registered", name)
	}

	if err := registerExchange(name, exchangeFn); err != nil {
		return err
	}

	exchange, err := exchangeFn(m.ctx, m.clientPool, bindManager, m.opts.GetExchange(name))
	if err != nil {
		return err
	}

	m.exchanges[name] = exchange
	return nil
}

// GetExchange retrieves an exchange.
func (m *Manager) GetExchange(name utils.ExchangeType) (Exchange, bool) {
	exchange, ok := m.exchanges[name]
	return exchange, ok
}

// GetExchanges retrieves the exchanges map.
func (m *Manager) GetExchanges() map[utils.ExchangeType]Exchange {
	return m.exchanges
}
