package observers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/unpackdev/solgo/bytecode"
	"github.com/unpackdev/solgo/contracts"
	"github.com/unpackdev/solgo/utils"
)

type BlockEntry struct {
	// UUID is never set by the observer and can be used by you to implement relational database and other
	// features if you wish to. Basically so you are aware of the block UUID when processing transactions as example...
	UUID      uuid.UUID
	NetworkID utils.NetworkID
	Network   utils.Network
	Strategy  utils.Strategy
	Block     *types.Block
}

type TransactionEntry struct {
	UUID      uuid.UUID `json:"uuid"`
	BlockUUID uuid.UUID `json:"block_uuid"`

	NetworkID         utils.NetworkID                        `json:"network_id"`
	Network           utils.Network                          `json:"network"`
	Strategy          utils.Strategy                         `json:"strategy"`
	TransactionType   utils.TransactionMethodType            `json:"transaction_type"`
	TransactionMethod *bytecode.Transaction                  `json:"transaction_method"`
	ContractAddress   common.Address                         `json:"contract_address"`
	Contract          *contracts.Contract                    `json:"contract"`
	Sender            common.Address                         `json:"sender"`
	SenderType        utils.AddressType                      `json:"sender_type"`
	SenderContract    *contracts.Contract                    `json:"sender_contract"`
	Recipient         common.Address                         `json:"recipient"`
	RecipientType     utils.AddressType                      `json:"recipient_type"`
	RecipientContract *contracts.Contract                    `json:"recipient_contract"`
	BlockHeader       *types.Header                          `json:"block_header"`
	Transaction       *types.Transaction                     `json:"transaction"`
	Receipt           *types.Receipt                         `json:"receipt"`
	LogContracts      map[common.Address]*contracts.Contract `json:"log_contracts"`
	Logs              []*contracts.Log                       `json:"logs"`
}

type LogEntry struct {
	UUID              uuid.UUID `json:"uuid"`
	BlockUUID         uuid.UUID `json:"block_uuid"`
	TransactionUUID   uuid.UUID `json:"transaction_uuid"`
	NetworkID         utils.NetworkID
	Network           utils.Network
	Strategy          utils.Strategy
	TransactionType   utils.TransactionMethodType
	TransactionMethod *bytecode.Transaction
	ContractAddress   common.Address
	Contract          *contracts.Contract
	Sender            common.Address
	SenderType        utils.AddressType
	SenderContract    *contracts.Contract
	Recipient         common.Address
	RecipientType     utils.AddressType
	RecipientContract *contracts.Contract
	BlockHeader       *types.Header
	Transaction       *types.Transaction
	Receipt           *types.Receipt
	LogContract       *contracts.Contract
	Log               *contracts.Log
}

type ContractEntry struct {
	UUID              uuid.UUID
	NetworkID         utils.NetworkID
	BlockUUID         uuid.UUID `json:"block_uuid"`
	TransactionUUID   uuid.UUID `json:"transaction_uuid"`
	Network           utils.Network
	Strategy          utils.Strategy
	TransactionType   utils.TransactionMethodType
	TransactionMethod *bytecode.Transaction
	ContractAddress   common.Address
	Contract          *contracts.Contract
	Sender            common.Address
	SenderType        utils.AddressType
	SenderContract    *contracts.Contract
	Recipient         common.Address
	RecipientType     utils.AddressType
	RecipientContract *contracts.Contract
	BlockHeader       *types.Header
	Transaction       *types.Transaction
	Receipt           *types.Receipt
}

// IsToken checks if contract entry contains token information.
func (c *ContractEntry) IsToken() bool {
	if c.Contract != nil {
		return c.Contract.IsToken()
	}

	return false
}
