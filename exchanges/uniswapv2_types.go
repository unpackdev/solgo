package exchanges

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/utils"
	"github.com/unpackdev/solgo/utils/entities"
)

type UniswapV2PairReserves struct {
	Token0    common.Address `json:"token0"`
	Token1    common.Address `json:"token1"`
	Reserve0  *big.Int       `json:"reserve0"`
	Reserve1  *big.Int       `json:"reserve1"`
	BlockTime time.Time      `json:"block_time"`
}

type AuditApprovalResults struct {
	Detected          bool            `json:"detected"`
	ApprovalRequested bool            `json:"approval_requested"`
	Approved          bool            `json:"approved"`
	TxHash            common.Hash     `json:"transaction_hash"`
	Receipt           bool            `json:"receipt"`
	ReceiptStatus     uint64          `json:"receipt_status"`
	GasUsedRaw        uint64          `json:"gas_used_raw"`
	GasUsed           string          `json:"gas_used"`
	Logs              []*bytecode.Log `json:"logs"`
}

type AuditSwapResults struct {
	Detected              bool             `json:"detected"`
	Failure               bool             `json:"failure"`
	SwapRequested         bool             `json:"swap_requested"`
	PairDetails           []common.Address `json:"pair_details"`
	TxHash                common.Hash      `json:"transaction_hash"`
	Receipt               bool             `json:"receipt"`
	ReceiptStatus         uint64           `json:"receipt_status"`
	Logs                  []*bytecode.Log  `json:"logs"`
	GasUsedRaw            uint64           `json:"gas_used_raw"`
	GasUsed               string           `json:"gas_used"`
	SwapReceivedAmountRaw *big.Int         `json:"swap_received_amount_raw"`
	SwapReceivedAmount    string           `json:"swap_received_amount"`
	ReceivedAmountRaw     *big.Int         `json:"received_amount_raw"`
	ReceivedAmount        string           `json:"received_amount"`
	TaxRaw                *big.Int         `json:"tax_raw"`
	Tax                   string           `json:"tax"`
}

type AuditBuyOrSellResults struct {
	Detected bool                  `json:"detected"`
	Approval *AuditApprovalResults `json:"approval"`
	Results  *AuditSwapResults     `json:"results"`
}

type UniswapV2TradeDescriptor struct {
	Network              utils.Network          `json:"network"`
	NetworkID            utils.NetworkID        `json:"network_id"`
	Simulation           bool                   `json:"simulation"`
	ExchangeType         utils.ExchangeType     `json:"exchange_type"`
	SpenderAddress       common.Address         `json:"spender_address"`
	SpenderBalanceBefore *big.Int               `json:"spender_balance_before"`
	SpenderBalanceAfter  *big.Int               `json:"spender_balance_after"`
	AmountRaw            *big.Int               `json:"amount_raw"`
	Amount               string                 `json:"amount"`
	RouterAddress        common.Address         `json:"router_address"`
	FactoryAddress       common.Address         `json:"factory_address"`
	WETHAddress          common.Address         `json:"weth_address"`
	PairAddress          common.Address         `json:"pair_address"`
	PairReserves         *UniswapV2PairReserves `json:"pair_reserves"`
	UsdToEthPriceRaw     *entities.Price        `json:"-"`
	UsdToEthPrice        string                 `json:"usd_to_eth_price"`
	EthToUsdPriceRaw     *entities.Price        `json:"-"`
	EthToUsdPrice        string                 `json:"eth_to_usd_price"`
	Price                *entities.Price        `json:"-"`
	PricePerToken        string                 `json:"price_per_token"`
	PricePerTokenUsdRaw  *entities.Price        `json:"-"`
	PricePerTokenUsd     string                 `json:"price_per_token_usd"`
	MaxAmountRaw         *big.Int               `json:"max_amount_raw"`
	MaxAmount            string                 `json:"max_amount"`
	Trade                *AuditBuyOrSellResults `json:"trade"`
}
