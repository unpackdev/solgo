package opcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpCodeFunctions(t *testing.T) {
	tests := []struct {
		op      OpCode
		isPush  bool
		isJump  bool
		isArith bool
		isComp  bool
		isBit   bool
		isBlock bool
		isStack bool
		isMem   bool
		isStore bool
		isFlow  bool
		isSys   bool
		isSelfD bool
	}{
		{PUSH1, true, false, false, false, false, false, false, false, false, false, false, false},
		{JUMP, false, true, false, false, false, false, true, false, false, true, false, false},
		{ADD, false, false, true, false, false, false, false, false, false, false, false, false},
		{LT, false, false, false, true, false, false, false, false, false, false, false, false},
		{AND, false, false, false, false, true, false, false, false, false, false, false, false},
		{BLOCKHASH, false, false, false, false, false, true, false, false, false, false, false, false},
		{POP, false, false, false, false, false, false, true, false, false, false, false, false},
		{MLOAD, false, false, false, false, false, false, true, true, false, false, false, false},
		{SLOAD, false, false, false, false, false, false, true, false, true, false, false, false},
		{STOP, false, false, false, false, false, false, false, false, false, true, false, false},
		{CREATE, false, false, false, false, false, false, false, false, false, false, true, false},
		{JUMPI, false, true, false, false, false, false, true, false, false, true, false, false},
		{MUL, false, false, true, false, false, false, false, false, false, false, false, false},
		{GT, false, false, false, true, false, false, false, false, false, false, false, false},
		{OR, false, false, false, false, true, false, false, false, false, false, false, false},
		{COINBASE, false, false, false, false, false, true, false, false, false, false, false, false},
		{MSTORE, false, false, false, false, false, false, true, true, false, false, false, false},
		{SSTORE, false, false, false, false, false, false, true, false, true, false, false, false},
		{RETURN, false, false, false, false, false, false, false, false, false, true, true, false},
		{CALL, false, false, false, false, false, false, false, false, false, false, true, false},
		{EQ, false, false, false, true, false, false, false, false, false, false, false, false},
		{BYTE, false, false, false, false, true, false, false, false, false, false, false, false},
		{TIMESTAMP, false, false, false, false, false, true, false, false, false, false, false, false},
		{SUB, false, false, true, false, false, false, false, false, false, false, false, false},
		{DIV, false, false, true, false, false, false, false, false, false, false, false, false},
		{SDIV, false, false, true, false, false, false, false, false, false, false, false, false},
		{MOD, false, false, true, false, false, false, false, false, false, false, false, false},
		{SMOD, false, false, true, false, false, false, false, false, false, false, false, false},
		{ADDMOD, false, false, true, false, false, false, false, false, false, false, false, false},
		{MULMOD, false, false, true, false, false, false, false, false, false, false, false, false},
		{EXP, false, false, true, false, false, false, false, false, false, false, false, false},
		{SIGNEXTEND, false, false, true, false, false, false, false, false, false, false, false, false},
		{SLT, false, false, false, true, false, false, false, false, false, false, false, false},
		{SGT, false, false, false, true, false, false, false, false, false, false, false, false},
		{ISZERO, false, false, false, true, false, false, false, false, false, false, false, false},
		{XOR, false, false, false, false, true, false, false, false, false, false, false, false},
		{NOT, false, false, false, false, true, false, false, false, false, false, false, false},
		{SHL, false, false, false, false, true, false, false, false, false, false, false, false},
		{SHR, false, false, false, false, true, false, false, false, false, false, false, false},
		{SAR, false, false, false, false, true, false, false, false, false, false, false, false},
		{DIFFICULTY, false, false, false, false, false, true, false, false, false, false, false, false},
		{GASLIMIT, false, false, false, false, false, true, false, false, false, false, false, false},
		{PC, false, false, false, false, false, false, true, false, false, true, false, false},
		{MSIZE, false, false, false, false, false, false, true, true, false, true, false, false},
		{GAS, false, false, false, false, false, false, true, false, false, true, false, false},
		{REVERT, false, false, false, false, false, false, false, false, false, true, true, false},
		{INVALID, false, false, false, false, false, false, false, false, false, true, true, false},
		{SELFDESTRUCT, false, false, false, false, false, false, false, false, false, true, true, true},
		{CALLCODE, false, false, false, false, false, false, false, false, false, false, true, false},
		{DELEGATECALL, false, false, false, false, false, false, false, false, false, false, true, false},
		{CREATE2, false, false, false, false, false, false, false, false, false, false, true, false},
		{STATICCALL, false, false, false, false, false, false, false, false, false, false, true, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.isPush, tt.op.IsPush(), "IsPush mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isJump, tt.op.IsJump(), "IsJump mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isArith, tt.op.IsArithmetic(), "IsArithmetic mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isComp, tt.op.IsComparison(), "IsComparison mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isBit, tt.op.IsBitwise(), "IsBitwise mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isBlock, tt.op.IsBlockInformation(), "IsBlockInformation mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isStack, tt.op.IsStack(), "IsStack mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isMem, tt.op.IsMemory(), "IsMemory mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isStore, tt.op.IsStorage(), "IsStorage mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isFlow, tt.op.IsFlowControl(), "IsFlowControl mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isSys, tt.op.IsSystem(), "IsSystem mismatch for opcode %v", tt.op)
		assert.Equal(t, tt.isSelfD, tt.op.IsSelfDestruct(), "IsSelfDestruct mismatch for opcode %v", tt.op)
	}
}
