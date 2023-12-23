package bytecode

import (
	"bytes"
	"fmt"
	"strings"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/unpackdev/solgo/utils"
)

// Transaction represents a decoded Ethereum transaction, including its ABI, signature, and method information.
type Transaction struct {
	Abi            string                      `json:"abi"`             // ABI of the transaction
	SignatureBytes []byte                      `json:"signature_bytes"` // Raw signature of the transaction
	Signature      string                      `json:"signature"`       // Signature of the transaction
	Type           utils.TransactionMethodType `json:"type"`            // Type of the transaction
	Name           string                      `json:"name"`            // Name of the transaction
	Method         *abi.Method                 `json:"-"`
	Inputs         map[string]interface{}      `json:"inputs"` // List of arguments in the transaction
}

// DecodeTransactionFromAbi decodes an Ethereum transaction from its raw data and ABI.
// It extracts the method signature and arguments, and constructs a Transaction object
// containing this information along with the ABI for the method.
//
// data is the raw transaction data. abiData is the ABI of the smart contract in JSON format.
// This function returns a pointer to a Transaction object or an error if the decoding fails.
func DecodeTransactionFromAbi(data []byte, abiData []byte) (*Transaction, error) {
	// The first 4 bytes of the data represent the ID of the method in the ABI.
	methodSigData := data[:4]

	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		return nil, fmt.Errorf("failed to parse abi: %s", err)
	}

	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		return nil, fmt.Errorf("failed to get method by id: %s", err)
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return nil, fmt.Errorf("failed to unpack inputs into map: %s", err)
	}

	abi, _ := utils.MethodToABI(method)

	return &Transaction{
		Abi:            abi,
		SignatureBytes: methodSigData,
		Signature:      method.String(),
		Name:           method.Name,
		Method:         method,
		Type:           utils.TransactionMethodType(strings.ToLower(method.Name)),
		Inputs:         inputsMap,
	}, nil
}
