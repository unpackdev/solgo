package tokens

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (t *Token) CalculateTotalBurnedSupply(ctx context.Context) error {
	burnAddresses := []common.Address{
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		common.HexToAddress("0x000000000000000000000000000000000000dead"),
		common.HexToAddress("0x000000000000000000000000000000000000dEaD"),
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		common.HexToAddress("0x0000000000000000000000000000000000000003"),
		common.HexToAddress("0x0000000000000000000000000000000000000004"),
		common.HexToAddress("0x0000000000000000000000000000000000000005"),
		common.HexToAddress("0x0000000000000000000000000000000000000006"),
		common.HexToAddress("0x0000000000000000000000000000000000000007"),
		common.HexToAddress("0x0000000000000000000000000000000000000008"),
		common.HexToAddress("0x0000000000000000000000000000000000000009"),
	}

	totalBurned := big.NewInt(0)

	for _, address := range burnAddresses {
		balance, err := t.ResolveBalance(ctx, t.descriptor.Address, t.tokenBind, address)
		if err != nil {
			return fmt.Errorf("failed to resolve token total burned supply for %s : %s", address.Hex(), err)
		}

		totalBurned.Add(totalBurned, balance)
	}

	t.descriptor.TotalBurnedSupply = totalBurned
	return nil
}
