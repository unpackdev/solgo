package opcode

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// MatchFunctionSignature checks if a given function signature matches any of the decompiled instructions.
func (d *Decompiler) MatchFunctionSignature(signature string) bool {
	// Remove "0x" prefix if present
	signature = strings.TrimPrefix(signature, "0x")

	for _, instruction := range d.instructions {
		if instruction.OpCode == CALL && len(instruction.Args) >= 4 {
			functionSig := common.Bytes2Hex(instruction.Args[:4])
			if functionSig == signature {
				return true
			}
		}
	}
	return false
}

func (d *Decompiler) MatchInstruction(instruction Instruction) bool {
	for _, inst := range d.instructions {
		if inst.Offset == instruction.Offset && inst.OpCode == instruction.OpCode && bytes.Equal(inst.Args, instruction.Args) {
			return true
		}
	}
	return false
}

// GetStateVariables returns the instruction trees for all state variables declared in the bytecode.
func (d *Decompiler) GetStateVariables() ([]*InstructionTree, error) {
	// Initialize slice to hold state variable trees with estimated capacity
	stateVariables := make([]*InstructionTree, 0)

	// Iterate through instructions to find state variable declarations
	for _, instruction := range d.instructions {
		// Check if the instruction is a state variable declaration
		if d.isStateVariableDeclaration(instruction) {
			// Build instruction tree for the state variable declaration
			tree := NewInstructionTree(instruction)
			stateVariables = append(stateVariables, tree)
		}
	}

	return stateVariables, nil
}

// Function to determine if an instruction declares a state variable
func (d *Decompiler) isStateVariableDeclaration(instruction Instruction) bool {
	// Check if the instruction is a PUSH operation followed by a SSTORE operation
	if instruction.OpCode.IsPush() && instruction.Offset+1 < len(d.instructions) {
		nextInstruction := d.instructions[instruction.Offset+1]
		if nextInstruction.OpCode == SSTORE {
			return true
		}
	}
	return false
}

// GetEvents returns the instruction trees for all events declared in the bytecode.
func (d *Decompiler) GetEvents() ([]*InstructionTree, error) {
	// Initialize slice to hold event trees with estimated capacity
	events := make([]*InstructionTree, 0)

	// Iterate through instructions to find event definitions
	for _, instruction := range d.instructions {
		// Check if the instruction is an event definition
		if isEventDefinition(instruction) {
			// Parse event arguments
			args, err := parseEventArguments(instruction)
			if err != nil {
				return nil, fmt.Errorf("error parsing event arguments: %v", err)
			}

			// Build instruction tree for the event definition with arguments
			tree := NewInstructionTree(instruction)
			tree.Instruction.Args = args // Append event arguments
			events = append(events, tree)
		}
	}

	return events, nil
}

// Function to determine if an instruction declares an event
func isEventDefinition(instruction Instruction) bool {
	// Check if the instruction is a LOG operation
	return instruction.OpCode == LOG0 || instruction.OpCode == LOG1 ||
		instruction.OpCode == LOG2 || instruction.OpCode == LOG3 || instruction.OpCode == LOG4
}

// Function to parse event arguments from instruction
func parseEventArguments(instruction Instruction) ([]byte, error) {
	// Check if the instruction has arguments
	if len(instruction.Args) > 0 {
		// For LOG0, no additional argument is needed
		if instruction.OpCode == LOG0 {
			return []byte{}, nil
		}

		// For other LOG operations, extract event arguments
		eventArgs := instruction.Args[32:] // Skip the first 32 bytes (topic)
		return eventArgs, nil
	}
	return nil, nil
}
