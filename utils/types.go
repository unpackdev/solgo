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

	UnknownLogEventType LogEventType = "unknown"

	AnvilSimulator SimulatorType = "anvil"

	SimulatorAccountType AccountType = "simulator"
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
