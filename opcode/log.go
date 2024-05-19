package opcode

import (
	"github.com/ethereum/go-ethereum/common"
)

// DecodeLOG1 returns all instructions with the LOG1 OpCode and decodes their arguments.
func (d *Decompiler) DecodeLOG1() []Instruction {
	return d.decodeLogInstructions(LOG1, 1)
}

// DecodeLOG2 returns all instructions with the LOG2 OpCode and decodes their arguments.
func (d *Decompiler) DecodeLOG2() []Instruction {
	return d.decodeLogInstructions(LOG2, 2)
}

// DecodeLOG3 returns all instructions with the LOG3 OpCode and decodes their arguments.
func (d *Decompiler) DecodeLOG3() []Instruction {
	return d.decodeLogInstructions(LOG3, 3)
}

// DecodeLOG4 returns all instructions with the LOG4 OpCode and decodes their arguments.
func (d *Decompiler) DecodeLOG4() []Instruction {
	return d.decodeLogInstructions(LOG4, 4)
}

// decodeLogInstructions processes the LOG instructions and decodes their arguments.
func (d *Decompiler) decodeLogInstructions(opCode OpCode, topicCount int) []Instruction {
	instructions := d.GetInstructionsByOpCode(opCode)

	for _, instruction := range instructions {
		data, topics := d.decodeLogArgs(instruction.Offset, topicCount)
		instruction.Data = data
		_ = topics
	}

	return instructions
}

// decodeLogArgs decodes the arguments for a LOG instruction.
func (d *Decompiler) decodeLogArgs(offset, topicCount int) ([]byte, []string) {
	var stack [][]byte
	currentIndex := -1

	// Find the index of the LOG instruction
	for i, instr := range d.instructions {
		if instr.Offset == offset {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return nil, nil
	}

	// Process instructions to reconstruct the stack state
	for i := 0; i < currentIndex; i++ {
		instr := d.instructions[i]
		switch {
		case instr.OpCode.IsPush():
			stack = append(stack, instr.Args)
		case instr.OpCode == SWAP1:
			if len(stack) >= 2 {
				stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
			}
		case instr.OpCode == SWAP2:
			if len(stack) >= 3 {
				stack[len(stack)-1], stack[len(stack)-3] = stack[len(stack)-3], stack[len(stack)-1]
			}
		case instr.OpCode == SWAP3:
			if len(stack) >= 4 {
				stack[len(stack)-1], stack[len(stack)-4] = stack[len(stack)-4], stack[len(stack)-1]
			}
		case instr.OpCode == SWAP4:
			if len(stack) >= 5 {
				stack[len(stack)-1], stack[len(stack)-5] = stack[len(stack)-5], stack[len(stack)-1]
			}
		case instr.OpCode == DUP1:
			if len(stack) >= 1 {
				stack = append(stack, stack[len(stack)-1])
			}
		case instr.OpCode == DUP2:
			if len(stack) >= 2 {
				stack = append(stack, stack[len(stack)-2])
			}
		case instr.OpCode == DUP3:
			if len(stack) >= 3 {
				stack = append(stack, stack[len(stack)-3])
			}
		case instr.OpCode == DUP4:
			if len(stack) >= 4 {
				stack = append(stack, stack[len(stack)-4])
			}
		case instr.OpCode == POP:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	// Collect the data and topics for the LOG instruction
	if len(stack) < topicCount+1 {
		return nil, nil
	}

	topics := make([]string, 0)
	var data []byte
	topicCountFound := 0

	// Iterate from the end of the stack to find topics and data
	for i := len(stack) - 1; i >= 0; i-- {
		if len(stack[i]) == 32 && topicCountFound < topicCount {
			topics = append([]string{common.Bytes2Hex(stack[i])}, topics...) // Prepend to maintain order
			topicCountFound++
		} else if data == nil {
			data = stack[i]
		}

		// Break if we have found all topics and data
		if topicCountFound == topicCount && data != nil {
			break
		}
	}

	if data == nil {
		return nil, nil
	}

	return data, topics
}
