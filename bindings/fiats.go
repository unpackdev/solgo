package bindings

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

const (
	USDC BindingType = "USDC"
)

type Fiat struct {
	*Manager
	ctx  context.Context
	opts []*BindOptions
}

func NewFiat(ctx context.Context, manager *Manager, opts []*BindOptions) (*Fiat, error) {
	if opts == nil {
		opts = DefaultPinksaleBindOptions()
	}

	for _, opt := range opts {
		if err := opt.Validate(); err != nil {
			return nil, err
		}
	}

	// Now lets register all the bindings with the manager
	for _, opt := range opts {
		for _, network := range opt.Networks {
			if _, err := manager.RegisterBinding(network, opt.NetworkID, opt.Type, opt.Address, opt.ABI); err != nil {
				return nil, err
			}
		}
	}

	return &Fiat{
		Manager: manager,
		ctx:     ctx,
		opts:    opts,
	}, nil
}

// WatchLockEvents starts watching for Lock events.
func (ps *Fiat) WatchLockEvents(ctx context.Context) error {
	return nil
}

// WatchUnlockEvents starts watching for Unlock events.
func (ps *Fiat) WatchUnlockEvents(ctx context.Context) error {
	return nil
}

func DefaultFiatBindOptions() []*BindOptions {
	return []*BindOptions{
		{
			Networks:  []utils.Network{utils.Ethereum, utils.AnvilNetwork},
			NetworkID: utils.EthereumNetworkID,
			Type:      USDC,
			Address:   common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			ABI:       `[{"constant":false,"inputs":[{"name":"newImplementation","type":"address"}],"name":"upgradeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newImplementation","type":"address"},{"name":"data","type":"bytes"}],"name":"upgradeToAndCall","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"implementation","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newAdmin","type":"address"}],"name":"changeAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_implementation","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":false,"name":"previousAdmin","type":"address"},{"indexed":false,"name":"newAdmin","type":"address"}],"name":"AdminChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"implementation","type":"address"}],"name":"Upgraded","type":"event"}]`,
		},
	}
}
