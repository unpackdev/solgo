package simulator

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/clients"
)

type Simulator struct {
	ctx              context.Context
	clientPool       *clients.ClientPool
	opts             *Options
	simulatedBackend *backends.SimulatedBackend
	faucets          []*Faucet
}

func NewSimulator(ctx context.Context, clientPool *clients.ClientPool, opts *Options) (*Simulator, error) {
	if opts == nil {
		opts = NewDefaultOptions() // Use default options if none provided
	}

	s := &Simulator{
		ctx:        ctx,
		clientPool: clientPool,
		opts:       opts,
	}

	err := s.initSimulatedBackend()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Simulator) initSimulatedBackend() error {
	// Initialize the allocation with pre-defined addresses
	alloc := core.GenesisAlloc{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
	}

	// Initialize faucets and add them to the allocation
	for i := 0; i < s.opts.NumberOfFaucets; i++ {
		faucet, err := NewFaucet()
		if err != nil {
			return err
		}
		s.faucets = append(s.faucets, faucet)
		alloc[faucet.Address] = core.GenesisAccount{
			Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9)),
		}
	}

	// Initialize the simulated backend with the specified gas limit
	s.simulatedBackend = backends.NewSimulatedBackend(alloc, s.opts.GasLimit)

	return nil
}

func (s *Simulator) GetChainID() *big.Int {
	return s.simulatedBackend.Blockchain().Config().ChainID
}

func (s *Simulator) Faucets() []*Faucet {
	return s.faucets
}

func (s *Simulator) Faucet(address common.Address) *Faucet {
	for _, faucet := range s.faucets {
		if faucet.Address == address {
			return faucet
		}
	}
	return nil
}

func (s *Simulator) GetBackend() *backends.SimulatedBackend {
	return s.simulatedBackend
}

func (s *Simulator) DeployContract(auth *bind.TransactOpts, abiData string, bytecode []byte, constructorParams ...interface{}) (common.Address, *types.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		return common.Address{}, nil, err
	}

	address, tx, _, err := bind.DeployContract(auth, parsedABI, bytecode, s.simulatedBackend, constructorParams...)
	if err != nil {
		return common.Address{}, tx, err
	}
	s.simulatedBackend.Commit()

	return address, tx, nil
}
