package opcode

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	opcode_pb "github.com/txpull/protos/dist/go/opcode"
)

// Decompiler is responsible for decompiling Ethereum bytecode into a set of instructions.
type Decompiler struct {
	ctx          context.Context // The context for the decompiler.
	bytecode     []byte          // The bytecode to be decompiled.
	bytecodeSize uint64          // The size of the bytecode.
	instructions []Instruction   // The resulting set of instructions after decompilation.
}

// NewDecompiler initializes a new Decompiler with the given bytecode.
func NewDecompiler(ctx context.Context, b []byte) (*Decompiler, error) {
	return &Decompiler{
		ctx:          ctx,
		bytecode:     b,
		bytecodeSize: uint64(len(b)),
		instructions: []Instruction{},
	}, nil
}

// GetBytecode returns the bytecode associated with the Decompiler.
func (d *Decompiler) GetBytecode() []byte {
	return d.bytecode
}

// GetBytecodeSize returns the size of the bytecode.
func (d *Decompiler) GetBytecodeSize() uint64 {
	return d.bytecodeSize
}

// Decompile processes the bytecode and populates the instructions slice.
func (d *Decompiler) Decompile() error {
	if d.bytecodeSize < 1 {
		return ErrEmptyBytecode
	}

	offset := 0
	for offset < len(d.bytecode) {
		op := OpCode(d.bytecode[offset])
		instruction := Instruction{
			Offset:      offset,
			OpCode:      op,
			Args:        []byte{},
			Description: op.GetDescription(),
		}

		if op.IsPush() {
			argSize := int(op) - int(PUSH1) + 1
			if offset+argSize >= len(d.bytecode) {
				break
			}
			instruction.Args = d.bytecode[offset+1 : offset+argSize+1]
			offset += argSize
		}

		d.instructions = append(d.instructions, instruction)
		offset++
	}
	return nil
}

// GetInstructionsByOpCode returns all instructions that match the given OpCode.
func (d *Decompiler) GetInstructionsByOpCode(op OpCode) []Instruction {
	var callInstructions []Instruction
	for _, instruction := range d.instructions {
		if instruction.OpCode == op {
			callInstructions = append(callInstructions, instruction)
		}
	}
	return callInstructions
}

// GetInstructions returns all decompiled instructions.
func (d *Decompiler) GetInstructions() []Instruction {
	return d.instructions
}

// IsOpCode checks if the given instruction matches the provided OpCode.
func (d *Decompiler) IsOpCode(instruction Instruction, op OpCode) bool {
	return instruction.OpCode == op
}

// OpCodeFound checks if the given OpCode exists in the decompiled instructions.
func (d *Decompiler) OpCodeFound(op OpCode) bool {
	for _, instruction := range d.instructions {
		if instruction.OpCode == op {
			return true
		}
	}
	return false
}

// ToProto converts the decompiled instructions to a protobuf representation.
func (d *Decompiler) ToProto() *opcode_pb.Root {
	instructions := make([]*opcode_pb.Instruction, 0)
	for _, instruction := range d.instructions {
		instructions = append(instructions, instruction.ToProto())
	}

	return &opcode_pb.Root{
		Instructions: instructions,
	}
}

// String provides a string representation of the decompiled instructions.
func (d *Decompiler) String() string {
	var buf bytes.Buffer

	for _, instr := range d.instructions {
		offset := fmt.Sprintf("0x%04x", instr.Offset)
		opCode := instr.OpCode.String()

		buf.WriteString(offset + " " + opCode)

		if len(instr.Args) > 0 {
			buf.WriteString(" " + common.Bytes2Hex(instr.Args))
		}

		desc := instr.OpCode.GetDescription()
		if desc != "" {
			buf.WriteString(" // " + desc)
		}

		buf.WriteString("\n")
	}

	return buf.String()
}

// GetInstructionTreeFormatted returns a formatted string representation of the opcode execution tree
// starting from the provided instruction. The output is indented based on the provided indent string.
func (d *Decompiler) GetInstructionTreeFormatted(instruction Instruction, indent string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s0x%04x %s\n", indent, instruction.Offset, instruction.OpCode.String()))

	childIndent := indent + "   "
	for _, child := range d.GetChildrenByOffset(instruction.Offset) {
		builder.WriteString(d.GetInstructionTreeFormatted(child, childIndent))
	}

	return builder.String()
}

// GetChildrenByOffset retrieves a slice of Instructions that are immediate children (subsequent instructions)
// of the instruction at the provided offset.
func (d *Decompiler) GetChildrenByOffset(offset int) []Instruction {
	var children []Instruction
	for _, instr := range d.instructions {
		if instr.Offset == offset+1 {
			children = append(children, instr)
		}
	}
	return children
}
