package exchanges

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/utils"
)

type ExchangeOptions struct {
	Network        utils.Network  `json:"network" yaml:"network" mapstructure:"network"`
	Exchange       ExchangeType   `json:"exchange" yaml:"exchange" mapstructure:"exchange"`
	RouterAddress  common.Address `json:"router_address" yaml:"router_address" mapstructure:"router_address"`
	FactoryAddress common.Address `json:"factory_address" yaml:"factory_address" mapstructure:"factory_address"`
}

func (o *ExchangeOptions) Validate() error {
	if o.Network == "" {
		return fmt.Errorf("network cannot be empty")
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

type Options struct {
	Exchanges []ExchangeOptions `json:"exchanges" yaml:"exchanges" mapstructure:"exchanges"`
}

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

func (o *Options) GetExchange(name ExchangeType) *ExchangeOptions {
	for _, exchange := range o.Exchanges {
		if exchange.Exchange == name {
			return &exchange
		}
	}

	return nil
}

func DefaultOptions() *Options {
	return &Options{
		Exchanges: []ExchangeOptions{
			{
				Network:  utils.Ethereum,
				Exchange: UniswapV2,
				// https://etherscan.io/address/0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
				RouterAddress: common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"),
				// https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f
				FactoryAddress: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			},
			{
				Network:  utils.Ethereum,
				Exchange: UniswapV3,
				// https://etherscan.io/address/0xe592427a0aece92de3edee1f18e0157c05861564
				RouterAddress: common.HexToAddress("0xe592427a0aece92de3edee1f18e0157c05861564"),
				// https://etherscan.io/address/0x1F98431c8aD98523631AE4a59f267346ea31F984
				FactoryAddress: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
			},
			{
				Network: utils.Ethereum,
				// https://docs.sushi.com/docs/Products/Classic%20AMM/Overview
				Exchange: Sushiswap,
				// https://etherscan.io/address/0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F
				RouterAddress: common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"),
				// https://etherscan.io/address/0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac
				FactoryAddress: common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"),
			},
			{
				Network: utils.Bsc,
				// https://docs.pancakeswap.finance/developers/smart-contracts/pancakeswap-exchange/v2-contracts
				Exchange: PancakeswapV2,
				// https://bscscan.com/address/0x10ed43c718714eb63d5aa57b78b54704e256024e
				RouterAddress: common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"),
				// https://bscscan.com/address/0xca143ce32fe78f1f7019d7d551a6402fc5350c73
				FactoryAddress: common.HexToAddress("0xca143ce32fe78f1f7019d7d551a6402fc5350c73"),
			},
			/* 			{
				Network: utils.Bsc,
				// https://docs.pancakeswap.finance/developers/smart-contracts/pancakeswap-exchange/v2-contracts
				Exchange: PancakeswapV3,
				// https://bscscan.com/address/0x1b81D678ffb9C0263b24A97847620C99d213eB14
				RouterAddress: common.HexToAddress("0x1b81D678ffb9C0263b24A97847620C99d213eB14"),
				// https://bscscan.com/address/0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865
				FactoryAddress: common.HexToAddress("0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"),
			}, */
		},
	}
}
