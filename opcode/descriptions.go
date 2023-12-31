package opcode

// descriptions maps each OpCode to its corresponding description.
var descriptions = map[OpCode]string{
	STOP:           "Halts execution.",
	ADD:            "Adds the top two stack items.",
	MUL:            "Multiplies the top two stack items.",
	SUB:            "Subtracts the second stack item from the first.",
	DIV:            "Divides the first stack item by the second.",
	SDIV:           "Signed division operation.",
	MOD:            "Modulus remainder operation.",
	SMOD:           "Signed modulus operation.",
	EXP:            "Exponential operation.",
	NOT:            "Bitwise NOT operation.",
	LT:             "Checks if the first item is less than the second.",
	GT:             "Checks if the first item is greater than the second.",
	SLT:            "Signed less than.",
	SGT:            "Signed greater than.",
	EQ:             "Checks if the two top stack items are equal.",
	ISZERO:         "Checks if the top stack item is zero.",
	SIGNEXTEND:     "Extend length of two's complement signed integer.",
	AND:            "Bitwise AND operation between the two top stack items.",
	OR:             "Bitwise OR operation between the two top stack items.",
	XOR:            "Bitwise XOR operation between the two top stack items.",
	BYTE:           "Retrieve single byte from word.",
	SHL:            "Shift left.",
	SHR:            "Shift right.",
	SAR:            "Arithmetic shift right.",
	ADDMOD:         "Modulo addition operation.",
	MULMOD:         "Modulo multiplication operation.",
	KECCAK256:      "Computes the Keccak-256 hash of input.",
	ADDRESS:        "Get address of currently executing account.",
	BALANCE:        "Get balance of the given account.",
	ORIGIN:         "Get execution origination address.",
	CALLER:         "Get caller address.",
	CALLVALUE:      "Get deposited value by the instruction/transaction responsible for this execution.",
	CALLDATALOAD:   "Get input data of current environment.",
	CALLDATASIZE:   "Get size of input data in current environment.",
	CALLDATACOPY:   "Copy input data in current environment to memory.",
	CHAINID:        "Get the chain ID of the current chain.",
	BASEFEE:        "Get the base fee of the current block.",
	BLOBHASH:       "Get the hash of the current blob.",
	DELEGATECALL:   "Message-call into this account with an alternative account’s code, but persisting the current values for `sender` and `value`.",
	STATICCALL:     "Static message-call into an account.",
	CODESIZE:       "Get size of code running in current environment.",
	CODECOPY:       "Copy code running in current environment to memory.",
	GASPRICE:       "Get price of gas in current environment.",
	EXTCODESIZE:    "Get size of an account's code.",
	EXTCODECOPY:    "Copy an account's code to memory.",
	RETURNDATASIZE: "Get size of output data from the previous instruction.",
	RETURNDATACOPY: "Copy output data from the previous instruction to memory.",
	EXTCODEHASH:    "Get the hash of an account's code.",
	BLOCKHASH:      "Get the hash of one of the 256 most recent complete blocks.",
	COINBASE:       "Get the block's beneficiary address.",
	TIMESTAMP:      "Get the block's timestamp.",
	NUMBER:         "Get the block's number.",
	DIFFICULTY:     "Get the block's difficulty.",
	GASLIMIT:       "Get the block's gas limit.",
	SELFBALANCE:    "Get balance of the current account.",
	POP:            "Remove item from stack.",
	MLOAD:          "Load word from memory.",
	MSTORE:         "Save word to memory.",
	MSTORE8:        "Save byte to memory.",
	SLOAD:          "Load word from storage.",
	SSTORE:         "Save word to storage.",
	JUMP:           "Alter the program counter.",
	JUMPI:          "Conditionally alter the program counter.",
	PC:             "Get the value of the program counter prior to the increment.",
	MSIZE:          "Get size of active memory in bytes.",
	GAS:            "Get the amount of available gas, including the corresponding reduction the amount of available gas.",
	JUMPDEST:       "Mark a valid destination for jumps.",
	PUSH0:          "Push 0 bytes onto the stack.",
	PUSH1:          "Push 1 byte onto the stack.",
	PUSH2:          "Push 2 bytes onto the stack.",
	PUSH3:          "Push 3 bytes onto the stack.",
	PUSH4:          "Push 4 bytes onto the stack.",
	PUSH5:          "Push 5 bytes onto the stack.",
	PUSH6:          "Push 6 bytes onto the stack.",
	PUSH7:          "Push 7 bytes onto the stack.",
	PUSH8:          "Push 8 bytes onto the stack.",
	PUSH9:          "Push 9 bytes onto the stack.",
	PUSH10:         "Push 10 bytes onto the stack.",
	PUSH11:         "Push 11 bytes onto the stack.",
	PUSH12:         "Push 12 bytes onto the stack.",
	PUSH13:         "Push 13 bytes onto the stack.",
	PUSH14:         "Push 14 bytes onto the stack.",
	PUSH15:         "Push 15 bytes onto the stack.",
	PUSH16:         "Push 16 bytes onto the stack.",
	PUSH17:         "Push 17 bytes onto the stack.",
	PUSH18:         "Push 18 bytes onto the stack.",
	PUSH19:         "Push 19 bytes onto the stack.",
	PUSH20:         "Push 20 bytes onto the stack.",
	PUSH21:         "Push 21 bytes onto the stack.",
	PUSH22:         "Push 22 bytes onto the stack.",
	PUSH23:         "Push 23 bytes onto the stack.",
	PUSH24:         "Push 24 bytes onto the stack.",
	PUSH25:         "Push 25 bytes onto the stack.",
	PUSH26:         "Push 26 bytes onto the stack.",
	PUSH27:         "Push 27 bytes onto the stack.",
	PUSH28:         "Push 28 bytes onto the stack.",
	PUSH29:         "Push 29 bytes onto the stack.",
	PUSH30:         "Push 30 bytes onto the stack.",
	PUSH31:         "Push 31 bytes onto the stack.",
	PUSH32:         "Push 32 bytes onto the stack.",
	DUP1:           "Duplicates the 1st stack item.",
	DUP2:           "Duplicates the 2nd stack item.",
	DUP3:           "Duplicates the 3rd stack item.",
	DUP4:           "Duplicates the 4th stack item.",
	DUP5:           "Duplicates the 5th stack item.",
	DUP6:           "Duplicates the 6th stack item.",
	DUP7:           "Duplicates the 7th stack item.",
	DUP8:           "Duplicates the 8th stack item.",
	DUP9:           "Duplicates the 9th stack item.",
	DUP10:          "Duplicates the 10th stack item.",
	DUP11:          "Duplicates the 11th stack item.",
	DUP12:          "Duplicates the 12th stack item.",
	DUP13:          "Duplicates the 13th stack item.",
	DUP14:          "Duplicates the 14th stack item.",
	DUP15:          "Duplicates the 15th stack item.",
	DUP16:          "Duplicates the 16th stack item.",
	SWAP1:          "Swaps the top stack item with the 2nd stack item.",
	SWAP2:          "Swaps the top stack item with the 3rd stack item.",
	SWAP3:          "Swaps the top stack item with the 4th stack item.",
	SWAP4:          "Swaps the top stack item with the 5th stack item.",
	SWAP5:          "Swaps the top stack item with the 6th stack item.",
	SWAP6:          "Swaps the top stack item with the 7th stack item.",
	SWAP7:          "Swaps the top stack item with the 8th stack item.",
	SWAP8:          "Swaps the top stack item with the 9th stack item.",
	SWAP9:          "Swaps the top stack item with the 10th stack item.",
	SWAP10:         "Swaps the top stack item with the 11th stack item.",
	SWAP11:         "Swaps the top stack item with the 12th stack item.",
	SWAP12:         "Swaps the top stack item with the 13th stack item.",
	SWAP13:         "Swaps the top stack item with the 14th stack item.",
	SWAP14:         "Swaps the top stack item with the 15th stack item.",
	SWAP15:         "Swaps the top stack item with the 16th stack item.",
	SWAP16:         "Swaps the top stack item with the 17th stack item.",
	LOG0:           "Appends log record with no topics.",
	LOG1:           "Appends log record with 1 topic.",
	LOG2:           "Appends log record with 2 topics.",
	LOG3:           "Appends log record with 3 topics.",
	LOG4:           "Appends log record with 4 topics.",
	TLOAD:          "Load from transaction context.",
	TSTORE:         "Store to transaction context.",
	CREATE:         "Create a new account with associated code.",
	CREATE2:        "Create a new account with associated code at a specific address.",
	CALL:           "Message-call into an account.",
	RETURN:         "Halt execution returning output data.",
	CALLCODE:       "Message-call into this account with another account's code.",
	REVERT:         "Halt execution reverting state changes but returning data and remaining gas.",
	INVALID:        "Designated invalid instruction.",
	SELFDESTRUCT:   "Halt execution and register account for later deletion.",
}

// GetDescription retrieves the description of the OpCode.
// If the opcode is not found in the descriptions map, it returns an empty string.
func (op OpCode) GetDescription() string {
	if desc, exists := descriptions[op]; exists {
		return desc
	}
	return ""
}
