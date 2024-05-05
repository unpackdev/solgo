package opcode

// Arithmetic operations (0x0 range).
const (
	// STOP halts execution.
	STOP OpCode = 0x0
	// ADD performs addition operation.
	ADD OpCode = 0x1
	// MUL performs multiplication operation.
	MUL OpCode = 0x2
	// SUB performs subtraction operation.
	SUB OpCode = 0x3
	// DIV performs division operation.
	DIV OpCode = 0x4
	// SDIV performs signed division operation.
	SDIV OpCode = 0x5
	// MOD returns the remainder after division.
	MOD OpCode = 0x6
	// SMOD returns the signed remainder after division.
	SMOD OpCode = 0x7
	// ADDMOD performs addition followed by modulo operation.
	ADDMOD OpCode = 0x8
	// MULMOD performs multiplication followed by modulo operation.
	MULMOD OpCode = 0x9
	// EXP performs exponential operation.
	EXP OpCode = 0xa
	// SIGNEXTEND extends the sign bit to the left.
	SIGNEXTEND OpCode = 0xb
)

// Comparison operations (0x10 range).
const (
	// LT checks if one value is less than another.
	LT OpCode = 0x10
	// GT checks if one value is greater than another.
	GT OpCode = 0x11
	// SLT checks if one signed value is less than another.
	SLT OpCode = 0x12
	// SGT checks if one signed value is greater than another.
	SGT OpCode = 0x13
	// EQ checks if two values are equal.
	EQ OpCode = 0x14
	// ISZERO checks if a value is zero.
	ISZERO OpCode = 0x15
	// AND performs a bitwise AND operation.
	AND OpCode = 0x16
	// OR performs a bitwise OR operation.
	OR OpCode = 0x17
	// XOR performs a bitwise XOR operation.
	XOR OpCode = 0x18
	// NOT performs a bitwise NOT operation.
	NOT OpCode = 0x19
	// BYTE retrieves a specific byte from a value.
	BYTE OpCode = 0x1a
	// SHL shifts a value to the left.
	SHL OpCode = 0x1b
	// SHR shifts a value to the right.
	SHR OpCode = 0x1c
	// SAR performs an arithmetic right shift.
	SAR OpCode = 0x1d
)

// Cryptographic operations (0x20 range).
const (
	// KECCAK256 computes the Keccak-256 hash.
	KECCAK256 OpCode = 0x20
)

// Closure state operations (0x30 range).
const (
	// ADDRESS retrieves the address of the current contract.
	ADDRESS OpCode = 0x30
	// BALANCE retrieves the balance of a given address.
	BALANCE OpCode = 0x31
	// ORIGIN retrieves the address that originated the current call.
	ORIGIN OpCode = 0x32
	// CALLER retrieves the address of the caller.
	CALLER OpCode = 0x33
	// CALLVALUE retrieves the value sent with the call.
	CALLVALUE OpCode = 0x34
	// CALLDATALOAD loads data from the call payload.
	CALLDATALOAD OpCode = 0x35
	// CALLDATASIZE retrieves the size of the call data.
	CALLDATASIZE OpCode = 0x36
	// CALLDATACOPY copies call data to memory.
	CALLDATACOPY OpCode = 0x37
	// CODESIZE retrieves the size of the code.
	CODESIZE OpCode = 0x38
	// CODECOPY copies code to memory.
	CODECOPY OpCode = 0x39
	// GASPRICE retrieves the price of gas.
	GASPRICE OpCode = 0x3a
	// EXTCODESIZE retrieves the size of an external contract's code.
	EXTCODESIZE OpCode = 0x3b
	// EXTCODECOPY copies external contract code to memory.
	EXTCODECOPY OpCode = 0x3c
	// RETURNDATASIZE retrieves the size of the return data.
	RETURNDATASIZE OpCode = 0x3d
	// RETURNDATACOPY copies return data to memory.
	RETURNDATACOPY OpCode = 0x3e
	// EXTCODEHASH retrieves the hash of an external contract's code.
	EXTCODEHASH OpCode = 0x3f
)

// Block operations (0x40).
const (
	// BLOCKHASH retrieves the hash of a block.
	BLOCKHASH OpCode = 0x40
	// COINBASE retrieves the address of the block miner.
	COINBASE OpCode = 0x41
	// TIMESTAMP retrieves the timestamp of the block.
	TIMESTAMP OpCode = 0x42
	// NUMBER retrieves the block number.
	NUMBER OpCode = 0x43
	// DIFFICULTY retrieves the difficulty of the block.
	DIFFICULTY OpCode = 0x44
	// RANDOM retrieves a random value. Note: It has the same opcode as DIFFICULTY.
	RANDOM OpCode = 0x44 // Same as DIFFICULTY
	// PREVRANDAO retrieves the previous RANDAO value. Note: It has the same opcode as DIFFICULTY.
	PREVRANDAO OpCode = 0x44 // Same as DIFFICULTY
	// GASLIMIT retrieves the gas limit for the block.
	GASLIMIT OpCode = 0x45
	// CHAINID retrieves the chain ID.
	CHAINID OpCode = 0x46
	// SELFBALANCE retrieves the balance of the contract itself.
	SELFBALANCE OpCode = 0x47
	// BASEFEE retrieves the base fee for the block.
	BASEFEE OpCode = 0x48
	// BLOBHASH retrieves the blob hash.
	BLOBHASH OpCode = 0x49
)

// Storage and execution operations (0x50).
const (
	// POP removes the top item from the stack.
	POP OpCode = 0x50
	// MLOAD loads a word from memory at a specific address.
	MLOAD OpCode = 0x51
	// MSTORE stores a word to memory at a specific address.
	MSTORE OpCode = 0x52
	// MSTORE8 stores a byte to memory at a specific address.
	MSTORE8 OpCode = 0x53
	// SLOAD loads a word from storage at a specific address.
	SLOAD OpCode = 0x54
	// SSTORE stores a word to storage at a specific address.
	SSTORE OpCode = 0x55
	// JUMP sets the program counter to a specific address.
	JUMP OpCode = 0x56
	// JUMPI conditionally sets the program counter to a specific address.
	JUMPI OpCode = 0x57
	// PC retrieves the value of the program counter.
	PC OpCode = 0x58
	// MSIZE retrieves the size of active memory.
	MSIZE OpCode = 0x59
	// GAS retrieves the amount of gas remaining.
	GAS OpCode = 0x5a
	// JUMPDEST marks a valid destination for jumps.
	JUMPDEST OpCode = 0x5b
	// PUSH0 pushes a zero byte onto the stack.
	PUSH0 OpCode = 0x5f
)

// Push operations (0x60)
const (
	// PUSH1 pushes a 1-byte item onto the stack.
	PUSH1 OpCode = 0x60 + iota
	// PUSH2 pushes a 2-byte item onto the stack.
	PUSH2
	// PUSH3 pushes a 3-byte item onto the stack.
	PUSH3
	// PUSH4 pushes a 4-byte item onto the stack.
	PUSH4
	// PUSH5 pushes a 5-byte item onto the stack.
	PUSH5
	// PUSH6 pushes a 6-byte item onto the stack.
	PUSH6
	// PUSH7 pushes a 7-byte item onto the stack.
	PUSH7
	// PUSH8 pushes an 8-byte item onto the stack.
	PUSH8
	// PUSH9 pushes a 9-byte item onto the stack.
	PUSH9
	// PUSH10 pushes a 10-byte item onto the stack.
	PUSH10
	// PUSH11 pushes an 11-byte item onto the stack.
	PUSH11
	// PUSH12 pushes a 12-byte item onto the stack.
	PUSH12
	// PUSH13 pushes a 13-byte item onto the stack.
	PUSH13
	// PUSH14 pushes a 14-byte item onto the stack.
	PUSH14
	// PUSH15 pushes a 15-byte item onto the stack.
	PUSH15
	// PUSH16 pushes a 16-byte item onto the stack.
	PUSH16
	// PUSH17 pushes a 17-byte item onto the stack.
	PUSH17
	// PUSH18 pushes an 18-byte item onto the stack.
	PUSH18
	// PUSH19 pushes a 19-byte item onto the stack.
	PUSH19
	// PUSH20 pushes a 20-byte item onto the stack.
	PUSH20
	// PUSH21 pushes a 21-byte item onto the stack.
	PUSH21
	// PUSH22 pushes a 22-byte item onto the stack.
	PUSH22
	// PUSH23 pushes a 23-byte item onto the stack.
	PUSH23
	// PUSH24 pushes a 24-byte item onto the stack.
	PUSH24
	// PUSH25 pushes a 25-byte item onto the stack.
	PUSH25
	// PUSH26 pushes a 26-byte item onto the stack.
	PUSH26
	// PUSH27 pushes a 27-byte item onto the stack.
	PUSH27
	// PUSH28 pushes a 28-byte item onto the stack.
	PUSH28
	// PUSH29 pushes a 29-byte item onto the stack.
	PUSH29
	// PUSH30 pushes a 30-byte item onto the stack.
	PUSH30
	// PUSH31 pushes a 31-byte item onto the stack.
	PUSH31
	// PUSH32 pushes a 32-byte item onto the stack.
	PUSH32
)

// Duplication operations (0x80)
const (
	// DUP1 duplicates the 1st item from the top of the stack.
	DUP1 = 0x80 + iota
	// DUP2 duplicates the 2nd item from the top of the stack.
	DUP2
	// DUP3 duplicates the 3rd item from the top of the stack.
	DUP3
	// DUP4 duplicates the 4th item from the top of the stack.
	DUP4
	// DUP5 duplicates the 5th item from the top of the stack.
	DUP5
	// DUP6 duplicates the 6th item from the top of the stack.
	DUP6
	// DUP7 duplicates the 7th item from the top of the stack.
	DUP7
	// DUP8 duplicates the 8th item from the top of the stack.
	DUP8
	// DUP9 duplicates the 9th item from the top of the stack.
	DUP9
	// DUP10 duplicates the 10th item from the top of the stack.
	DUP10
	// DUP11 duplicates the 11th item from the top of the stack.
	DUP11
	// DUP12 duplicates the 12th item from the top of the stack.
	DUP12
	// DUP13 duplicates the 13th item from the top of the stack.
	DUP13
	// DUP14 duplicates the 14th item from the top of the stack.
	DUP14
	// DUP15 duplicates the 15th item from the top of the stack.
	DUP15
	// DUP16 duplicates the 16th item from the top of the stack.
	DUP16
)

// Swap operations (0x90)
const (
	// SWAP1 swaps the top two items on the stack.
	SWAP1 = 0x90 + iota
	// SWAP2 swaps the top item with the 2nd item below it.
	SWAP2
	// SWAP3 swaps the top item with the 3rd item below it.
	SWAP3
	// SWAP4 swaps the top item with the 4th item below it.
	SWAP4
	// SWAP5 swaps the top item with the 5th item below it.
	SWAP5
	// SWAP6 swaps the top item with the 6th item below it.
	SWAP6
	// SWAP7 swaps the top item with the 7th item below it.
	SWAP7
	// SWAP8 swaps the top item with the 8th item below it.
	SWAP8
	// SWAP9 swaps the top item with the 9th item below it.
	SWAP9
	// SWAP10 swaps the top item with the 10th item below it.
	SWAP10
	// SWAP11 swaps the top item with the 11th item below it.
	SWAP11
	// SWAP12 swaps the top item with the 12th item below it.
	SWAP12
	// SWAP13 swaps the top item with the 13th item below it.
	SWAP13
	// SWAP14 swaps the top item with the 14th item below it.
	SWAP14
	// SWAP15 swaps the top item with the 15th item below it.
	SWAP15
	// SWAP16 swaps the top item with the 16th item below it.
	SWAP16
)

// Logging operations (0xa0)
const (
	// LOG0 logs a message with 0 indexed topics.
	LOG0 OpCode = 0xa0 + iota
	// LOG1 logs a message with 1 indexed topic.
	LOG1
	// LOG2 logs a message with 2 indexed topics.
	LOG2
	// LOG3 logs a message with 3 indexed topics.
	LOG3
	// LOG4 logs a message with 4 indexed topics.
	LOG4
)

// 0xb0 range.
const (
	// TLOAD loads from transactional storage.
	TLOAD OpCode = 0xb3
	// TSTORE stores to transactional storage.
	TSTORE OpCode = 0xb4
)

// Closure operations (0xf0)
const (
	// CREATE creates a new contract.
	CREATE OpCode = 0xf0
	// CALL initiates a new message call.
	CALL OpCode = 0xf1
	// CALLCODE calls a contract with a different code.
	CALLCODE OpCode = 0xf2
	// RETURN halts execution and returns output data.
	RETURN OpCode = 0xf3
	// DELEGATECALL calls a contract as a delegate.
	DELEGATECALL OpCode = 0xf4
	// CREATE2 creates a new contract with a deterministic address.
	CREATE2 OpCode = 0xf5

	// STATICCALL calls a contract without state modification.
	STATICCALL OpCode = 0xfa
	// REVERT stops execution and reverts state changes.
	REVERT OpCode = 0xfd
	// INVALID represents an invalid opcode.
	INVALID OpCode = 0xfe
	// SELFDESTRUCT destroys the current contract.
	SELFDESTRUCT OpCode = 0xff
)

// Since the opcodes aren't all in order we can't use a regular slice.
var opCodeToString = map[OpCode]string{
	// 0x0 range - arithmetic ops.
	STOP:       "STOP",
	ADD:        "ADD",
	MUL:        "MUL",
	SUB:        "SUB",
	DIV:        "DIV",
	SDIV:       "SDIV",
	MOD:        "MOD",
	SMOD:       "SMOD",
	EXP:        "EXP",
	NOT:        "NOT",
	LT:         "LT",
	GT:         "GT",
	SLT:        "SLT",
	SGT:        "SGT",
	EQ:         "EQ",
	ISZERO:     "ISZERO",
	SIGNEXTEND: "SIGNEXTEND",

	// 0x10 range - bit ops.
	AND:    "AND",
	OR:     "OR",
	XOR:    "XOR",
	BYTE:   "BYTE",
	SHL:    "SHL",
	SHR:    "SHR",
	SAR:    "SAR",
	ADDMOD: "ADDMOD",
	MULMOD: "MULMOD",

	// 0x20 range - crypto.
	KECCAK256: "KECCAK256",

	// 0x30 range - closure state.
	ADDRESS:        "ADDRESS",
	BALANCE:        "BALANCE",
	ORIGIN:         "ORIGIN",
	CALLER:         "CALLER",
	CALLVALUE:      "CALLVALUE",
	CALLDATALOAD:   "CALLDATALOAD",
	CALLDATASIZE:   "CALLDATASIZE",
	CALLDATACOPY:   "CALLDATACOPY",
	CODESIZE:       "CODESIZE",
	CODECOPY:       "CODECOPY",
	GASPRICE:       "GASPRICE",
	EXTCODESIZE:    "EXTCODESIZE",
	EXTCODECOPY:    "EXTCODECOPY",
	RETURNDATASIZE: "RETURNDATASIZE",
	RETURNDATACOPY: "RETURNDATACOPY",
	EXTCODEHASH:    "EXTCODEHASH",

	// 0x40 range - block operations.
	BLOCKHASH:   "BLOCKHASH",
	COINBASE:    "COINBASE",
	TIMESTAMP:   "TIMESTAMP",
	NUMBER:      "NUMBER",
	DIFFICULTY:  "DIFFICULTY",
	GASLIMIT:    "GASLIMIT",
	CHAINID:     "CHAINID",
	SELFBALANCE: "SELFBALANCE",
	BASEFEE:     "BASEFEE",
	BLOBHASH:    "BLOBHASH",

	// 0x50 range - 'storage' and execution.
	POP:      "POP",
	MLOAD:    "MLOAD",
	MSTORE:   "MSTORE",
	MSTORE8:  "MSTORE8",
	SLOAD:    "SLOAD",
	SSTORE:   "SSTORE",
	JUMP:     "JUMP",
	JUMPI:    "JUMPI",
	PC:       "PC",
	MSIZE:    "MSIZE",
	GAS:      "GAS",
	JUMPDEST: "JUMPDEST",
	PUSH0:    "PUSH0",

	// 0x60 range - pushes.
	PUSH1:  "PUSH1",
	PUSH2:  "PUSH2",
	PUSH3:  "PUSH3",
	PUSH4:  "PUSH4",
	PUSH5:  "PUSH5",
	PUSH6:  "PUSH6",
	PUSH7:  "PUSH7",
	PUSH8:  "PUSH8",
	PUSH9:  "PUSH9",
	PUSH10: "PUSH10",
	PUSH11: "PUSH11",
	PUSH12: "PUSH12",
	PUSH13: "PUSH13",
	PUSH14: "PUSH14",
	PUSH15: "PUSH15",
	PUSH16: "PUSH16",
	PUSH17: "PUSH17",
	PUSH18: "PUSH18",
	PUSH19: "PUSH19",
	PUSH20: "PUSH20",
	PUSH21: "PUSH21",
	PUSH22: "PUSH22",
	PUSH23: "PUSH23",
	PUSH24: "PUSH24",
	PUSH25: "PUSH25",
	PUSH26: "PUSH26",
	PUSH27: "PUSH27",
	PUSH28: "PUSH28",
	PUSH29: "PUSH29",
	PUSH30: "PUSH30",
	PUSH31: "PUSH31",
	PUSH32: "PUSH32",

	// 0x80 - dups.
	DUP1:  "DUP1",
	DUP2:  "DUP2",
	DUP3:  "DUP3",
	DUP4:  "DUP4",
	DUP5:  "DUP5",
	DUP6:  "DUP6",
	DUP7:  "DUP7",
	DUP8:  "DUP8",
	DUP9:  "DUP9",
	DUP10: "DUP10",
	DUP11: "DUP11",
	DUP12: "DUP12",
	DUP13: "DUP13",
	DUP14: "DUP14",
	DUP15: "DUP15",
	DUP16: "DUP16",

	// 0x90 - swaps.
	SWAP1:  "SWAP1",
	SWAP2:  "SWAP2",
	SWAP3:  "SWAP3",
	SWAP4:  "SWAP4",
	SWAP5:  "SWAP5",
	SWAP6:  "SWAP6",
	SWAP7:  "SWAP7",
	SWAP8:  "SWAP8",
	SWAP9:  "SWAP9",
	SWAP10: "SWAP10",
	SWAP11: "SWAP11",
	SWAP12: "SWAP12",
	SWAP13: "SWAP13",
	SWAP14: "SWAP14",
	SWAP15: "SWAP15",
	SWAP16: "SWAP16",

	// 0xa0 range - logging ops.
	LOG0: "LOG0",
	LOG1: "LOG1",
	LOG2: "LOG2",
	LOG3: "LOG3",
	LOG4: "LOG4",

	// 0xb0 range.
	TLOAD:  "TLOAD",
	TSTORE: "TSTORE",

	// 0xf0 range - closures.
	CREATE:       "CREATE",
	CALL:         "CALL",
	RETURN:       "RETURN",
	CALLCODE:     "CALLCODE",
	DELEGATECALL: "DELEGATECALL",
	CREATE2:      "CREATE2",
	STATICCALL:   "STATICCALL",
	REVERT:       "REVERT",
	INVALID:      "INVALID",
	SELFDESTRUCT: "SELFDESTRUCT",
}

var stringToOp = map[string]OpCode{
	"STOP":           STOP,
	"ADD":            ADD,
	"MUL":            MUL,
	"SUB":            SUB,
	"DIV":            DIV,
	"SDIV":           SDIV,
	"MOD":            MOD,
	"SMOD":           SMOD,
	"EXP":            EXP,
	"NOT":            NOT,
	"LT":             LT,
	"GT":             GT,
	"SLT":            SLT,
	"SGT":            SGT,
	"EQ":             EQ,
	"ISZERO":         ISZERO,
	"SIGNEXTEND":     SIGNEXTEND,
	"AND":            AND,
	"OR":             OR,
	"XOR":            XOR,
	"BYTE":           BYTE,
	"SHL":            SHL,
	"SHR":            SHR,
	"SAR":            SAR,
	"ADDMOD":         ADDMOD,
	"MULMOD":         MULMOD,
	"KECCAK256":      KECCAK256,
	"ADDRESS":        ADDRESS,
	"BALANCE":        BALANCE,
	"ORIGIN":         ORIGIN,
	"CALLER":         CALLER,
	"CALLVALUE":      CALLVALUE,
	"CALLDATALOAD":   CALLDATALOAD,
	"CALLDATASIZE":   CALLDATASIZE,
	"CALLDATACOPY":   CALLDATACOPY,
	"CHAINID":        CHAINID,
	"BASEFEE":        BASEFEE,
	"BLOBHASH":       BLOBHASH,
	"DELEGATECALL":   DELEGATECALL,
	"STATICCALL":     STATICCALL,
	"CODESIZE":       CODESIZE,
	"CODECOPY":       CODECOPY,
	"GASPRICE":       GASPRICE,
	"EXTCODESIZE":    EXTCODESIZE,
	"EXTCODECOPY":    EXTCODECOPY,
	"RETURNDATASIZE": RETURNDATASIZE,
	"RETURNDATACOPY": RETURNDATACOPY,
	"EXTCODEHASH":    EXTCODEHASH,
	"BLOCKHASH":      BLOCKHASH,
	"COINBASE":       COINBASE,
	"TIMESTAMP":      TIMESTAMP,
	"NUMBER":         NUMBER,
	"DIFFICULTY":     DIFFICULTY,
	"GASLIMIT":       GASLIMIT,
	"SELFBALANCE":    SELFBALANCE,
	"POP":            POP,
	"MLOAD":          MLOAD,
	"MSTORE":         MSTORE,
	"MSTORE8":        MSTORE8,
	"SLOAD":          SLOAD,
	"SSTORE":         SSTORE,
	"JUMP":           JUMP,
	"JUMPI":          JUMPI,
	"PC":             PC,
	"MSIZE":          MSIZE,
	"GAS":            GAS,
	"JUMPDEST":       JUMPDEST,
	"PUSH0":          PUSH0,
	"PUSH1":          PUSH1,
	"PUSH2":          PUSH2,
	"PUSH3":          PUSH3,
	"PUSH4":          PUSH4,
	"PUSH5":          PUSH5,
	"PUSH6":          PUSH6,
	"PUSH7":          PUSH7,
	"PUSH8":          PUSH8,
	"PUSH9":          PUSH9,
	"PUSH10":         PUSH10,
	"PUSH11":         PUSH11,
	"PUSH12":         PUSH12,
	"PUSH13":         PUSH13,
	"PUSH14":         PUSH14,
	"PUSH15":         PUSH15,
	"PUSH16":         PUSH16,
	"PUSH17":         PUSH17,
	"PUSH18":         PUSH18,
	"PUSH19":         PUSH19,
	"PUSH20":         PUSH20,
	"PUSH21":         PUSH21,
	"PUSH22":         PUSH22,
	"PUSH23":         PUSH23,
	"PUSH24":         PUSH24,
	"PUSH25":         PUSH25,
	"PUSH26":         PUSH26,
	"PUSH27":         PUSH27,
	"PUSH28":         PUSH28,
	"PUSH29":         PUSH29,
	"PUSH30":         PUSH30,
	"PUSH31":         PUSH31,
	"PUSH32":         PUSH32,
	"DUP1":           DUP1,
	"DUP2":           DUP2,
	"DUP3":           DUP3,
	"DUP4":           DUP4,
	"DUP5":           DUP5,
	"DUP6":           DUP6,
	"DUP7":           DUP7,
	"DUP8":           DUP8,
	"DUP9":           DUP9,
	"DUP10":          DUP10,
	"DUP11":          DUP11,
	"DUP12":          DUP12,
	"DUP13":          DUP13,
	"DUP14":          DUP14,
	"DUP15":          DUP15,
	"DUP16":          DUP16,
	"SWAP1":          SWAP1,
	"SWAP2":          SWAP2,
	"SWAP3":          SWAP3,
	"SWAP4":          SWAP4,
	"SWAP5":          SWAP5,
	"SWAP6":          SWAP6,
	"SWAP7":          SWAP7,
	"SWAP8":          SWAP8,
	"SWAP9":          SWAP9,
	"SWAP10":         SWAP10,
	"SWAP11":         SWAP11,
	"SWAP12":         SWAP12,
	"SWAP13":         SWAP13,
	"SWAP14":         SWAP14,
	"SWAP15":         SWAP15,
	"SWAP16":         SWAP16,
	"LOG0":           LOG0,
	"LOG1":           LOG1,
	"LOG2":           LOG2,
	"LOG3":           LOG3,
	"LOG4":           LOG4,
	"TLOAD":          TLOAD,
	"TSTORE":         TSTORE,
	"CREATE":         CREATE,
	"CREATE2":        CREATE2,
	"CALL":           CALL,
	"RETURN":         RETURN,
	"CALLCODE":       CALLCODE,
	"REVERT":         REVERT,
	"INVALID":        INVALID,
	"SELFDESTRUCT":   SELFDESTRUCT,
}

// Control flow opcodes
var controlFlowOpcodes = []OpCode{
	JUMP,
	JUMPI,
	JUMPDEST,
	RETURN,
	REVERT,
	INVALID,
	SELFDESTRUCT,
	CALL,
	CALLCODE,
	DELEGATECALL,
	CREATE,
	CREATE2,
	STOP,
}

// StringToOp finds the opcode whose name is stored in `str`.
func StringToOp(str string) OpCode {
	return stringToOp[str]
}
