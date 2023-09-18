package opcode

import (
	"fmt"

	opcode_pb "github.com/unpackdev/protos/dist/go/opcode"
)

// OpCode represents an Ethereum operation code (opcode).
// Opcodes are single byte values that represent a specific operation in the EVM.
// The OpCode type provides methods to retrieve the string representation of the opcode,
// and to determine if an opcode corresponds to specific operations like PUSH or JUMP.
type OpCode byte

// String returns the string representation of the OpCode.
// This method tries to fetch the human-readable name of the opcode from the opCodeToString map.
// If the opcode is not found in the map, it returns a formatted string indicating the undefined opcode.
func (op OpCode) String() string {
	str := opCodeToString[op]
	if len(str) == 0 {
		return fmt.Sprintf("opcode %#x not defined", int(op))
	}

	return str
}

// ToProto converts the opcode to a protobuf message.
func (op OpCode) ToProto() opcode_pb.OpCode {
	return opcode_pb.OpCode(op)
}

// IsPush checks if the given opcode is a PUSH opcode.
// In the Ethereum instruction set, there are several PUSH opcodes ranging from PUSH1 to PUSH32.
// These opcodes are used to place a series of bytes onto the stack.
func (op OpCode) IsPush() bool {
	return PUSH1 <= op && op <= PUSH32
}

// IsJump determines if an opcode corresponds to a jump operation.
// Jump operations in the EVM allow for altering the sequence of execution.
// This method checks for three specific jump-related opcodes: JUMP, JUMPI, and JUMPDEST.
func (op OpCode) IsJump() bool {
	switch op {
	case JUMP, JUMPI, JUMPDEST:
		return true
	default:
		return false
	}
}

// IsArithmetic checks if the given opcode corresponds to an arithmetic operation.
// Arithmetic operations in the EVM include addition, multiplication, subtraction, etc.
func (op OpCode) IsArithmetic() bool {
	switch op {
	case ADD, MUL, SUB, DIV, SDIV, MOD, SMOD, ADDMOD, MULMOD, EXP, SIGNEXTEND:
		return true
	default:
		return false
	}
}

// IsComparison checks if the given opcode corresponds to a comparison operation.
// Comparison operations in the EVM include less than, greater than, equal to, etc.
func (op OpCode) IsComparison() bool {
	switch op {
	case LT, GT, SLT, SGT, EQ, ISZERO:
		return true
	default:
		return false
	}
}

// IsBitwise checks if the given opcode corresponds to a bitwise operation.
// Bitwise operations in the EVM include AND, OR, XOR, NOT, etc.
func (op OpCode) IsBitwise() bool {
	switch op {
	case AND, OR, XOR, NOT, BYTE, SHL, SHR, SAR:
		return true
	default:
		return false
	}
}

// IsBlockInformation checks if the given opcode provides information about the current block.
// These opcodes provide details like the current block's hash, coinbase, timestamp, etc.
func (op OpCode) IsBlockInformation() bool {
	switch op {
	case BLOCKHASH, COINBASE, TIMESTAMP, NUMBER, DIFFICULTY, GASLIMIT:
		return true
	default:
		return false
	}
}

// IsStack checks if the given opcode is related to stack operations.
// Stack operations in the EVM include operations that interact with the main stack.
func (op OpCode) IsStack() bool {
	switch op {
	case POP, MLOAD, MSTORE, MSTORE8, SLOAD, SSTORE, JUMP, JUMPI, PC, MSIZE, GAS:
		return true
	default:
		return false
	}
}

// IsMemory checks if the given opcode is related to memory operations.
// Memory operations in the EVM include operations that interact with the memory segment.
func (op OpCode) IsMemory() bool {
	switch op {
	case MLOAD, MSTORE, MSTORE8, MSIZE:
		return true
	default:
		return false
	}
}

// IsStorage checks if the given opcode is related to storage operations.
// Storage operations in the EVM include operations that interact with the contract's storage.
func (op OpCode) IsStorage() bool {
	switch op {
	case SLOAD, SSTORE:
		return true
	default:
		return false
	}
}

// IsFlowControl checks if the given opcode is related to flow control operations.
// Flow control operations in the EVM include operations that alter the sequence of execution.
func (op OpCode) IsFlowControl() bool {
	switch op {
	case JUMP, JUMPI, PC, MSIZE, GAS, STOP, RETURN, REVERT, INVALID, SELFDESTRUCT:
		return true
	default:
		return false
	}
}

// IsSystem checks if the given opcode is a system operation.
// System operations in the EVM include operations like contract creation, external calls, etc.
func (op OpCode) IsSystem() bool {
	switch op {
	case CREATE, CALL, CALLCODE, RETURN, DELEGATECALL, CREATE2, STATICCALL, REVERT, INVALID, SELFDESTRUCT:
		return true
	default:
		return false
	}
}

// IsSelfDestruct checks if the given opcode corresponds to the SELFDESTRUCT operation.
// The SELFDESTRUCT opcode is used in the EVM to destroy the current contract, sending its funds to the provided address.
func (op OpCode) IsSelfDestruct() bool {
	return op == SELFDESTRUCT
}
