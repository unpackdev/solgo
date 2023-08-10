package opcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompiler_MatchFunctionSignature(t *testing.T) {
	// Create a new Decompiler instance and set the instructions.
	decompiler := &Decompiler{
		instructions: []Instruction{
			{
				OpCode: CALL,
				Args:   []byte{0x01, 0x02, 0x03, 0x04},
			},
			{
				OpCode: CALL,
				Args:   []byte{0x05, 0x06, 0x07, 0x08},
			},
		},
	}

	signature := "0x01020304"
	assert.True(t, decompiler.MatchFunctionSignature(signature))

	signature = "0x05060708"
	assert.True(t, decompiler.MatchFunctionSignature(signature))

	signature = "0x01020305" // Signature not present
	assert.False(t, decompiler.MatchFunctionSignature(signature))
}

func TestDecompiler_MatchInstruction(t *testing.T) {
	// Create a new Decompiler instance and set the instructions.
	decompiler := &Decompiler{
		instructions: []Instruction{
			{
				Offset: 0,
				OpCode: PUSH1,
				Args:   []byte{0x01},
			},
			{
				Offset: 1,
				OpCode: JUMP,
			},
			{
				Offset: 2,
				OpCode: ADD,
			},
		},
	}

	// Test matching of PUSH1 instruction
	instruction := Instruction{
		Offset: 0,
		OpCode: PUSH1,
		Args:   []byte{0x01},
	}
	matched := decompiler.MatchInstruction(instruction)
	assert.True(t, matched, "Expected PUSH1 instruction to match")

	// Test matching of JUMP instruction
	instruction = Instruction{
		Offset: 1,
		OpCode: JUMP,
	}
	matched = decompiler.MatchInstruction(instruction)
	assert.True(t, matched, "Expected JUMP instruction to match")

	// Test non-matching instruction
	instruction = Instruction{
		Offset: 2,
		OpCode: SSTORE,
	}
	matched = decompiler.MatchInstruction(instruction)
	assert.False(t, matched, "Expected SSTORE instruction to not match")
}
