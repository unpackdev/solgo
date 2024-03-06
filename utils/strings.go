package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

func StringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

func AddressInSlice(addr common.Address, list []common.Address) bool {
	for _, item := range list {
		if item == addr {
			return true
		}
	}

	return false
}

func NamedAddressInSlice(addr common.Address, list []NamedAddr) (*NamedAddr, bool) {
	for _, item := range list {
		if item.Addr == addr {
			return &item, true
		}
	}

	return nil, false
}

func ContainsBlacklistType(list []BlacklistType, item BlacklistType) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}

	return false
}

func ContainsUUID(list []uuid.UUID, item uuid.UUID) bool {
	for _, listItem := range list {
		if listItem.String() == item.String() {
			return true
		}
	}

	return false
}
