package bindings

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type BindOptions struct {
	Networks  []utils.Network
	NetworkID utils.NetworkID
	Name      string
	Type      BindingType
	Address   common.Address
	ABI       string
}

func (b *BindOptions) Validate() error {
	if len(b.Networks) == 0 {
		return fmt.Errorf("missing network")
	}
	if b.NetworkID == 0 {
		return fmt.Errorf("missing network id")
	}
	if b.Type == "" {
		return fmt.Errorf("missing binding type")
	}
	if b.Address == utils.ZeroAddress {
		return fmt.Errorf("missing address")
	}
	if b.ABI == "" {
		return fmt.Errorf("missing abi")
	}
	return nil
}
