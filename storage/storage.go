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
)

// Storage is a struct that encapsulates various components required to interact with Ethereum smart contracts.
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

// NewStorage creates a new Storage instance. It requires context, network configuration, and various components
// such as a client pool, simulator, etherscan provider, compiler, and binding manager.
// It returns an initialized Storage object or an error if the initialization fails.
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
		opts = NewDefaultOptions()
	}

	return &Storage{
		ctx:         ctx,
		network:     network,
		clientsPool: clientsPool,
		opts:        opts,
		simulator:   simulator,
		etherscan:   etherscan,
		compiler:    compiler,
		bindManager: bindManager,
	}, nil
}

// Describe queries and returns detailed information about a smart contract at a specific address.
// It requires context, the contract address, and the block number for the query.
// It returns a Reader instance containing contract details or an error if the query fails.
func (s *Storage) Describe(ctx context.Context, addr common.Address, atBlock *big.Int) (*Reader, error) {
	contract, err := contracts.NewContract(ctx, s.network, s.clientsPool, nil, nil, s.etherscan, s.compiler, s.bindManager, addr)
	if err != nil {
		return nil, err
	}

	decodedContract, err := s.DecodeContract(ctx, contract)
	if err != nil {
		return nil, err
	}

	return s._describe(ctx, addr, decodedContract, atBlock)
}

// DescribeFromContract is similar to Describe but starts with an existing contract instance.
// It requires context, a contract object, and the block number for the query.
// It returns a Reader instance containing contract details or an error if the query fails.
func (s *Storage) DescribeFromContract(ctx context.Context, contract *contracts.Contract, atBlock *big.Int) (*Reader, error) {
	if contract == nil {
		return nil, fmt.Errorf("contract is nil")
	}

	decodedContract, err := s.DecodeContract(ctx, contract)
	if err != nil {
		return nil, err
	}

	return s._describe(ctx, contract.GetAddress(), decodedContract, atBlock)
}

// _describe is an internal method that performs the actual description process of a contract.
// It constructs a Descriptor and utilizes a Reader to discover and calculate storage variables and layouts.
func (s *Storage) _describe(ctx context.Context, addr common.Address, contract *contracts.Contract, atBlock *big.Int) (*Reader, error) {
	descriptor := &Descriptor{
		Contract:         contract,
		Block:            atBlock,
		StateVariables:   make(map[string][]*Variable),
		TargetVariables:  make(map[string][]*Variable),
		ConstanVariables: make(map[string][]*Variable),
		StorageLayout:    make(map[string]*StorageLayout),
	}

	reader, err := NewReader(ctx, s, descriptor)
	if err != nil {
		return nil, err
	}

	if err := reader.DiscoverStorageVariables(); err != nil {
		return nil, err
	}

	if err := reader.CalculateStorageLayout(); err != nil {
		return nil, err
	}

	if err := s.populateStorageValues(ctx, addr, descriptor, atBlock); err != nil {
		return nil, err
	}

	return reader, nil
}

// populateStorageValues populates storage values for each slot in the descriptor's storage layout.
func (s *Storage) populateStorageValues(ctx context.Context, addr common.Address, descriptor *Descriptor, atBlock *big.Int) error {
	var lastStorageValue []byte
	var lastSlot int64 = -1
	var lastBlockNumber *big.Int

	for _, layout := range descriptor.GetStorageLayouts() {
		for _, slot := range layout.GetSlots() {
			if slot.Slot != lastSlot {
				blockNumber, storageValue, err := s.getStorageValueAt(ctx, addr, slot.Slot, atBlock)
				if err != nil {
					return err
				}
				slot.BlockNumber = blockNumber
				lastStorageValue = storageValue
				lastSlot = slot.Slot
				lastBlockNumber = blockNumber
			}

			slot.BlockNumber = lastBlockNumber
			slot.RawValue = lastStorageValue
			if err := convertStorageToValue(slot, lastStorageValue); err != nil {
				return err
			}
		}
	}

	return nil
}

// getStorageValueAt retrieves the storage value at a given slot for a contract.
func (s *Storage) getStorageValueAt(ctx context.Context, contractAddress common.Address, slot int64, blockNumber *big.Int) (*big.Int, []byte, error) {
	client := s.clientsPool.GetClientByGroup(s.network.String())
	if client == nil {
		return blockNumber, nil, fmt.Errorf("no client found for network %s", s.network)
	}

	if blockNumber == nil {
		latestHeader, err := client.BlockByNumber(ctx, nil)
		if err != nil {
			return blockNumber, nil, fmt.Errorf("failed to get latest block header: %v", err)
		}
		blockNumber = latestHeader.Number()
	}

	bigIntIndex := big.NewInt(slot)
	position := common.BigToHash(bigIntIndex)

	response, err := client.StorageAt(ctx, contractAddress, position, blockNumber)
	return blockNumber, response, err
}
