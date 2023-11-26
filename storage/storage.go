package storage

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0x19/solc-switch"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/clients"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/providers/etherscan"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
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

	layout, err := reader.GetStorageLayout()
	if err != nil {
		return nil, err
	}

	for contract, v := range layout {
		zap.L().Info(
			"Found storage slot",
			zap.String("contract_address", addr.Hex()),
			zap.String("contract_name", contract),
			zap.String("name", v.Name),
			zap.String("type", v.Type),
			zap.Int64("slot", v.Slot),
			zap.Any("offset", v.Offset),
			zap.Any("size", v.Size),
		)

		storageValue, err := s.StorageAt(ctx, addr, common.BytesToHash(v.Bytes()), atBlock)
		if err != nil {
			zap.L().Error(
				"Failed to get storage value",
				zap.Error(err),
				zap.String("contract_address", addr.Hex()),
				zap.String("contract_name", contract),
				zap.String("name", v.Name),
				zap.String("type", v.Type),
				zap.Int64("slot", v.Slot),
				zap.Any("offset", v.Offset),
				zap.Any("size", v.Size),
			)
		}

		zap.L().Info(
			"Storage value",
			zap.String("contract_address", addr.Hex()),
			zap.String("contract_name", contract),
			zap.String("name", v.Name),
			zap.String("type", v.Type),
			zap.Int64("slot", v.Slot),
			zap.Any("offset", v.Offset),
			zap.Any("size", v.Size),
			zap.Any("value", storageValue),
		)
	}

	return reader, nil
}

func (s *Storage) GetStorageDescriptorFromContract(ctx context.Context, c *contracts.Contract, atBlock *big.Int) (*Descriptor, error) {
	return &Descriptor{}, nil
}

func (s *Storage) StorageAt(ctx context.Context, contractAddress common.Address, hashedSlot common.Hash, blockNumber *big.Int) ([]byte, error) {
	client := s.clientsPool.GetClientByGroup(s.network.String())
	if client == nil {
		return nil, fmt.Errorf("no client found for network %s", s.network)
	}

	storageValue, err := client.StorageAt(ctx, contractAddress, hashedSlot, blockNumber)
	if err != nil {
		return nil, err
	}
	return storageValue, nil
}
