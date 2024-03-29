package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

// StringInSlice checks whether a string is present in a slice of strings.
func StringInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

// AddressInSlice checks whether an Ethereum address is present in a slice of addresses.
func AddressInSlice(addr common.Address, list []common.Address) bool {
	for _, item := range list {
		if item == addr {
			return true
		}
	}

	return false
}

// NamedAddressInSlice checks whether an Ethereum address is present in a slice
// of NamedAddr and returns the corresponding NamedAddr if found.
func NamedAddressInSlice(addr common.Address, list []NamedAddr) (*NamedAddr, bool) {
	for _, item := range list {
		if item.Addr == addr {
			return &item, true
		}
	}

	return nil, false
}

// ContainsBlacklistType checks whether a BlacklistType item is present in a slice of BlacklistType.
func ContainsBlacklistType(list []BlacklistType, item BlacklistType) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}

	return false
}

// ContainsTransactionMethodType checks whether a TransactionMethodType item is present
// in a slice of TransactionMethodType.
func ContainsTransactionMethodType(list []TransactionMethodType, item TransactionMethodType) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}

	return false
}

// ContainsUUID checks whether a UUID is present in a slice of UUIDs by comparing
// their string representations.
func ContainsUUID(list []uuid.UUID, item uuid.UUID) bool {
	for _, listItem := range list {
		if listItem.String() == item.String() {
			return true
		}
	}

	return false
}
