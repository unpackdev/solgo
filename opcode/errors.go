package opcode

import (
	"errors"
)

var (
	ErrEmptyBytecode = errors.New("bytecode is not set or empty bytecode provided")
)
