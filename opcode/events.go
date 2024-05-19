package opcode

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// EventTreeNode represents a node in the opcode execution tree that represents an event.
type EventTreeNode struct {
	EventSignatureHex string `json:"eventSignatureHex"`
	EventSignature    string `json:"event_signature"`
	EventBytesHex     string `json:"event_bytes_hex"`
	EventBytes        []byte `json:"event_bytes"`
	HasEventSignature bool   `json:"has_event_signature"`
	*TreeNode
}

// GetEvents iterates through LOG1 to LOG4 instructions, decodes their arguments, and collects them into EventTreeNode structures.
func (d *Decompiler) GetEvents() []*EventTreeNode {
	logInstructions := map[OpCode]int{
		LOG1: 1,
		LOG2: 2,
		LOG3: 3,
		LOG4: 4,
	}

	events := make([]*EventTreeNode, 0)

	for opCode, topicCount := range logInstructions {
		instructions := d.GetInstructionsByOpCode(opCode)
		for _, instruction := range instructions {
			_, topics := d.decodeLogArgs(instruction.Offset, topicCount)

			if len(topics) < 1 {
				continue
			}

			eventSignature := topics[0]
			eventSignatureBytes := common.Hex2Bytes(eventSignature)

			eventNode := &EventTreeNode{
				EventSignature:    eventSignature,
				EventSignatureHex: fmt.Sprintf("0x%s", eventSignature),
				EventBytesHex:     common.Bytes2Hex(eventSignatureBytes),
				EventBytes:        eventSignatureBytes,
				HasEventSignature: len(eventSignatureBytes) == 32,
				TreeNode:          &TreeNode{Instruction: instruction, Children: make([]*TreeNode, 0)},
			}

			events = append(events, eventNode)
		}
	}

	return events
}
