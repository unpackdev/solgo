package storage

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo"
	"github.com/unpackdev/solgo/utils"
)

type Storage struct {
	Network          utils.Network
	NetworkID        utils.NetworkID
	Address          common.Address
	Name             string
	CompilerVersion  utils.SemanticVersion
	ABI              string
	Optimized        bool
	OptimizationRuns int
	License          string
	EvmVersion       string
	Sources          *solgo.Sources
}

var storages = map[common.Address]*Storage{}

func GetStorage(addr common.Address) *Storage {
	return storages[addr]
}

func RegisterStorage(storage *Storage) error {
	if _, ok := storages[storage.Address]; ok {
		return fmt.Errorf("storage already registered %s - %s", storage.Address.Hex(), storage.Name)
	}

	storages[storage.Address] = storage
	return nil
}

func GetStorages() map[common.Address]*Storage {
	return storages
}
