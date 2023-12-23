package exchanges

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

// ExchangeOptions represents the configuration options for a single cryptocurrency exchange.
// It includes network settings, exchange type, and addresses for router and factory contracts.
type ExchangeOptions struct {
	Networks       []utils.Network    `json:"networks" yaml:"networks" mapstructure:"networks"`                      // List of networks that the exchange operates on.
	Exchange       utils.ExchangeType `json:"exchange" yaml:"exchange" mapstructure:"exchange"`                      // The type of exchange, e.g., UniswapV2, SushiSwap.
	RouterAddress  common.Address     `json:"router_address" yaml:"router_address" mapstructure:"router_address"`    // Ethereum address of the router contract.
	FactoryAddress common.Address     `json:"factory_address" yaml:"factory_address" mapstructure:"factory_address"` // Ethereum address of the factory contract.
}

// Validate checks the validity of the ExchangeOptions.
// It ensures that necessary fields are not empty and contain valid data.
func (o *ExchangeOptions) Validate() error {
	if len(o.Networks) < 1 {
		return fmt.Errorf("networks cannot be empty")
	}

	if o.Exchange == "" {
		return fmt.Errorf("exchange cannot be empty")
	}

	if !common.IsHexAddress(o.RouterAddress.Hex()) {
		return fmt.Errorf("router address cannot be empty")
	}

	if !common.IsHexAddress(o.FactoryAddress.Hex()) {
		return fmt.Errorf("factory address cannot be empty")
	}

	return nil
}

// Options represents the configuration for multiple exchanges.
type Options struct {
	Exchanges []ExchangeOptions `json:"exchanges" yaml:"exchanges" mapstructure:"exchanges"`
}

// Validate checks the validity of the Options.
// It ensures that there is at least one exchange specified and each exchange option is valid.
func (o *Options) Validate() error {
	if len(o.Exchanges) == 0 {
		return fmt.Errorf("you need to specify at least one exchange")
	}

	for _, exchange := range o.Exchanges {
		if err := exchange.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetExchange retrieves the ExchangeOptions for a specified exchange name.
// It returns nil if the exchange is not found in the options.
func (o *Options) GetExchange(name utils.ExchangeType) *ExchangeOptions {
	for _, exchange := range o.Exchanges {
		if exchange.Exchange == name {
			return &exchange
		}
	}

	return nil
}

// DefaultOptions provides a set of default options for common exchanges.
// This includes predefined settings for exchanges like Uniswap, SushiSwap, and Pancakeswap.
func DefaultOptions() *Options {
	return &Options{
		Exchanges: []ExchangeOptions{
			{
				Networks: []utils.Network{utils.Ethereum, utils.AnvilNetwork},
				Exchange: utils.UniswapV2,
				// https://etherscan.io/address/0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
				RouterAddress: common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"),
				// https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f
				FactoryAddress: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			},
			{
				Networks: []utils.Network{utils.Ethereum, utils.AnvilNetwork},
				Exchange: utils.UniswapV3,
				// https://etherscan.io/address/0xe592427a0aece92de3edee1f18e0157c05861564
				RouterAddress: common.HexToAddress("0xe592427a0aece92de3edee1f18e0157c05861564"),
				// https://etherscan.io/address/0x1F98431c8aD98523631AE4a59f267346ea31F984
				FactoryAddress: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
			},
			{
				Networks: []utils.Network{utils.Ethereum, utils.AnvilNetwork},
				// https://docs.sushi.com/docs/Products/Classic%20AMM/Overview
				Exchange: utils.SushiSwap,
				// https://etherscan.io/address/0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F
				RouterAddress: common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"),
				// https://etherscan.io/address/0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac
				FactoryAddress: common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"),
			},
			{
				Networks: []utils.Network{utils.Ethereum, utils.AnvilNetwork},
				// https://docs.pancakeswap.finance/developers/smart-contracts/pancakeswap-exchange/v2-contracts
				Exchange: utils.PancakeswapV2,
				// https://bscscan.com/address/0x10ed43c718714eb63d5aa57b78b54704e256024e
				RouterAddress: common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"),
				// https://bscscan.com/address/0xca143ce32fe78f1f7019d7d551a6402fc5350c73
				FactoryAddress: common.HexToAddress("0xca143ce32fe78f1f7019d7d551a6402fc5350c73"),
			},
		},
	}
}
