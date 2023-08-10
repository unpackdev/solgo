package opcode

import opcode_pb "github.com/txpull/protos/dist/go/opcode"

// Instruction represents an optcode instruction.
type Instruction struct {
	Offset      int    `json:"offset"`
	OpCode      OpCode `json:"opcode"`
	Args        []byte `json:"args"`
	Description string `json:"description"`
}

// GetOffset returns the offset of the instruction.
func (i Instruction) GetOffset() int {
	return i.Offset
}

// GetCode returns the opcode of the instruction.
func (i Instruction) GetCode() OpCode {
	return i.OpCode
}

// GetArgs returns the arguments of the instruction.
func (i Instruction) GetArgs() []byte {
	return i.Args
}

// GetDescription returns the description of the instruction.
func (i Instruction) GetDescription() string {
	return i.Description
}

// String returns the string representation of the instruction.
func (i Instruction) String() string {
	return i.OpCode.String()
}

// ToProto converts the instruction to a protobuf message.
func (i Instruction) ToProto() *opcode_pb.Instruction {
	return &opcode_pb.Instruction{
		Offset:      int32(i.Offset),
		OpCode:      i.OpCode.ToProto(),
		Args:        i.Args,
		Description: i.Description,
	}
}
