package utils

const (
	HeadStrategy    Strategy = "head"
	ArchiveStrategy Strategy = "archive"

	ZeroAddressRecipient AddressType = "zero_address"
	AddressRecipient     AddressType = "address"
	ContractRecipient    AddressType = "contract"

	UnknownTransactionMethodType TransactionMethodType = "unknown"
	ContractCreationType         TransactionMethodType = "contract_creation"
	TransferMethodType           TransactionMethodType = "transfer"

	UnknownLogEventType  LogEventType = "unknown"
	SwapLogEventType     LogEventType = "swap"
	TransferLogEventType LogEventType = "transfer"
	DepositLogEventType  LogEventType = "deposit"
	WithdrawLogEventType LogEventType = "withdraw"
	MintLogEventType     LogEventType = "mint"
	BurnLogEventType     LogEventType = "burn"

	NoSimulator    SimulatorType = "no_simulator"
	AnvilSimulator SimulatorType = "anvil"
	TraceSimulator SimulatorType = "trace"

	SimulatorAccountType AccountType = "simulator"
	SimpleAccountType    AccountType = "simple"
	KeystoreAccountType  AccountType = "keystore"

	NoExchange    ExchangeType = "no_exchange"
	UniswapV2     ExchangeType = "uniswap_v2"
	UniswapV3     ExchangeType = "uniswap_v3"
	SushiSwap     ExchangeType = "sushiswap"
	PancakeswapV2 ExchangeType = "pancakeswap_v2"

	Erc20TokenType  TokenType = "erc20"
	Erc721TokenType TokenType = "erc721"

	AntiWhalePinksale AntiWhaleType = "pinksale"
)

type Strategy string

func (s Strategy) String() string {
	return string(s)
}

type AddressType string

func (r AddressType) String() string {
	return string(r)
}

type TransactionMethodType string

func (t TransactionMethodType) String() string {
	return string(t)
}

type LogEventType string

func (t LogEventType) String() string {
	return string(t)
}

type AccountType string

func (t AccountType) String() string {
	return string(t)
}

type SimulatorType string

func (t SimulatorType) String() string {
	return string(t)
}

type ExchangeType string

func (t ExchangeType) String() string {
	return string(t)
}

type TokenType string

func (t TokenType) String() string {
	return string(t)
}

type AntiWhaleType string

func (t AntiWhaleType) String() string {
	return string(t)
}
