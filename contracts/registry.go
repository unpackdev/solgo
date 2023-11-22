package contracts

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

var contractsRegistry = map[utils.Network]map[common.Address]*Contract{}
var crMutex = sync.Mutex{}

func RegisterContract(network utils.Network, contract *Contract) {
	crMutex.Lock()
	defer crMutex.Unlock()
	if _, ok := contractsRegistry[network]; !ok {
		contractsRegistry[network] = map[common.Address]*Contract{}
	}

	contractsRegistry[network][contract.GetAddress()] = contract
}

func GetContract(network utils.Network, address common.Address) *Contract {
	crMutex.Lock()
	defer crMutex.Unlock()
	if _, ok := contractsRegistry[network]; !ok {
		return nil
	}

	return contractsRegistry[network][address]
}

func GetContracts(network utils.Network) map[common.Address]*Contract {
	crMutex.Lock()
	defer crMutex.Unlock()
	if _, ok := contractsRegistry[network]; !ok {
		return nil
	}

	return contractsRegistry[network]
}

func GetContractsRegistry() map[utils.Network]map[common.Address]*Contract {
	crMutex.Lock()
	defer crMutex.Unlock()
	return contractsRegistry
}
