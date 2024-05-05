package opcode

import "errors"

type InstructionTree struct {
	Instruction Instruction
	Children    []*InstructionTree
}

func NewInstructionTree(instruction Instruction) *InstructionTree {
	return &InstructionTree{
		Instruction: instruction,
		Children:    []*InstructionTree{},
	}
}

// Function to recursively build the instruction tree
func buildInstructionTree(instructions []Instruction, offset int) (*InstructionTree, int, error) {
	// Check if offset is out of range
	if offset < 0 || offset >= len(instructions) {
		return nil, offset, errors.New("offset out of range")
	}

	root := NewInstructionTree(instructions[offset])

	// Start from the next instruction after the root
	nextOffset := offset + 1

	// Iterate through the instructions to find child nodes
	for nextOffset < len(instructions) {
		inst := instructions[nextOffset]
		if isChildOf(root.Instruction, inst) {
			child, newOffset, err := buildInstructionTree(instructions, nextOffset)
			if err != nil {
				return nil, offset, err
			}
			root.Children = append(root.Children, child)
			nextOffset = newOffset
		} else {
			break
		}
	}

	return root, nextOffset, nil
}

// Function to determine if an instruction is a child of another instruction
func isChildOf(parent, child Instruction) bool {
	// Check if the child instruction's offset is greater than the parent instruction's offset
	if child.Offset <= parent.Offset {
		return false
	}

	// Check if the child instruction's offset falls within the range of the parent instruction
	if child.Offset >= parent.Offset && child.Offset < parent.Offset+len(parent.Args) {
		return true
	}

	// Check if the child instruction occurs within a basic block following the parent instruction
	for _, opcode := range controlFlowOpcodes {
		if child.OpCode == opcode {
			return true
		}
	}

	return false
}
