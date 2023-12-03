package bytecode

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/unpackdev/solgo/utils"
)

// Topic represents a decoded topic from an Ethereum event log.
type Topic struct {
	Name  string `json:"name"`  // Name of the topic
	Value any    `json:"value"` // Value of the topic, decoded into appropriate data type
}

// Log encapsulates details of a decoded Ethereum event log.
type Log struct {
	Event        *abi.Event         `json:"-"`             // Event is the ABI definition of the log's event
	Abi          string             `json:"abi"`           // Abi is the ABI string of the event
	SignatureHex common.Hash        `json:"signature_hex"` // SignatureHex is the hex-encoded signature of the event
	Signature    string             `json:"signature"`     // Signature of the event
	Type         utils.LogEventType `json:"type"`          // Type of the event as classified by solgo
	Name         string             `json:"name"`          // Name of the event
	Data         map[string]any     `json:"data"`          // Data contains the decoded event data
	Topics       []Topic            `json:"topics"`        // Topics are the decoded topics of the event
}

// DecodeLogFromAbi decodes an Ethereum event log using the provided ABI.
// It returns a structured Log containing the decoded data and topics.
func DecodeLogFromAbi(log *types.Log, abiData []byte) (*Log, error) {
	if log == nil || len(log.Topics) < 1 {
		return nil, fmt.Errorf("log is nil or has no topics")
	}

	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		return nil, fmt.Errorf("failed to parse abi: %s", err)
	}

	event, err := contractABI.EventByID(log.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get event by id %s: %s", log.Topics[0].Hex(), err)
	}

	data := make(map[string]any)
	if err := event.Inputs.UnpackIntoMap(data, log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack inputs into map: %s", err)
	}

	decodedTopics := make([]Topic, len(log.Topics))
	for i, topic := range log.Topics {
		if i == 0 {
			continue
		}

		decodedTopic, err := decodeTopic(topic, event.Inputs[i-1])
		if err != nil {
			return nil, fmt.Errorf("failed to decode topic: %s", err)
		}

		decodedTopics[i] = Topic{
			Name:  event.Inputs[i-1].Name,
			Value: decodedTopic,
		}
	}

	abi, _ := utils.EventToABI(event)

	toReturn := &Log{
		Event:        event,
		Abi:          abi,
		SignatureHex: log.Topics[0],
		Signature:    strings.TrimLeft(event.String(), "event "),
		Name:         event.Name,
		Type:         utils.LogEventType(strings.ToLower(event.Name)),
		Data:         data,
		Topics:       decodedTopics[1:], // Exclude the first topic (event signature)
	}

	return toReturn, nil
}

// decodeTopic decodes a single topic from an Ethereum event log based on its ABI argument type.
func decodeTopic(topic common.Hash, argument abi.Argument) (interface{}, error) {
	switch argument.Type.T {
	case abi.AddressTy:
		return common.BytesToAddress(topic.Bytes()), nil
	case abi.BoolTy:
		return topic[common.HashLength-1] == 1, nil
	case abi.IntTy, abi.UintTy:
		size := argument.Type.Size
		if size > 256 {
			return nil, fmt.Errorf("unsupported integer size: %d", size)
		}
		integer := new(big.Int).SetBytes(topic[:])
		if argument.Type.T == abi.IntTy && size < 256 {
			integer = adjustIntSize(integer, size)
		}
		return integer, nil
	case abi.StringTy:
		return topic, nil
	case abi.BytesTy, abi.FixedBytesTy:
		return topic.Bytes(), nil
	case abi.SliceTy, abi.ArrayTy:
		return nil, fmt.Errorf("array/slice decoding not implemented")
	default:
		return nil, fmt.Errorf("decoding for type %v not implemented", argument.Type.T)
	}
}

// adjustIntSize adjusts the size of an integer to match its ABI-specified size.
// This is particularly relevant for signed integers smaller than 256 bits.
func adjustIntSize(integer *big.Int, size int) *big.Int {
	if size == 256 || integer.Bit(size-1) == 0 {
		return integer
	}
	return new(big.Int).Sub(integer, new(big.Int).Lsh(big.NewInt(1), uint(size)))
}
