package storage

import (
	"context"
	"math/big"

	"github.com/0x19/solc-switch"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
)

type Storage struct {
	ctx         context.Context
	network     utils.Network
	clientsPool *clients.ClientPool
	opts        *Options
	simulator   *simulator.Simulator
	etherscan   *etherscan.EtherScanProvider
	compiler    *solc.Solc
	bindManager *bindings.Manager
}

func NewStorage(
	ctx context.Context,
	network utils.Network,
	clientsPool *clients.ClientPool,
	simulator *simulator.Simulator,
	etherscan *etherscan.EtherScanProvider,
	compiler *solc.Solc,
	bindManager *bindings.Manager,
	opts *Options,
) (*Storage, error) {
	if opts == nil {
		opts = NewDefaultOptions() // Use default options if none provided
	}

	s := &Storage{
		ctx:         ctx,
		network:     network,
		clientsPool: clientsPool,
		opts:        opts,
		simulator:   simulator,
		etherscan:   etherscan,
		compiler:    compiler,
		bindManager: bindManager,
	}

	return s, nil
}

func (s *Storage) GetStorageDescriptor(ctx context.Context, addr common.Address, atBlock *big.Int) (*Reader, error) {
	contract, err := contracts.NewContract(ctx, s.network, s.clientsPool, nil, nil, s.etherscan, s.compiler, s.bindManager, addr)
	if err != nil {
		return nil, err
	}

	if contract, err = s.DecodeContract(ctx, contract); err != nil {
		return nil, err
	}

	descriptor := &Descriptor{
		Contract:                     contract,
		Block:                        atBlock,
		StateVariables:               make(map[string][]*Variable),
		TargetVariables:              make(map[string][]*Variable),
		ConstantStorageSlotVariables: make(map[string][]*Variable),
	}
	reader, err := NewReader(ctx, s, descriptor)
	if err != nil {
		return nil, err
	}

	if err := reader.GetStorageVariables(); err != nil {
		return nil, err
	}

	if err := reader.GetStorageLayout(); err != nil {
		return nil, err
	}

	return reader, nil
}

func (s *Storage) GetStorageDescriptorFromContract(ctx context.Context, c *contracts.Contract, atBlock *big.Int) (*Descriptor, error) {
	return &Descriptor{}, nil
}
