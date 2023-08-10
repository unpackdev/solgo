package opcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	opcode_pb "github.com/txpull/protos/dist/go/opcode"
)

func TestInstructionMethods(t *testing.T) {
	tests := []struct {
		name        string
		instruction Instruction
		expected    struct {
			offset      int
			opCode      OpCode
			args        []byte
			description string
			stringRep   string
			protoRep    *opcode_pb.Instruction
		}
	}{
		{
			name: "Test Instruction Methods",
			instruction: Instruction{
				Offset:      1,
				OpCode:      STOP, // Assuming STOP is a valid OpCode
				Args:        []byte{0x01, 0x02},
				Description: "Test Description",
			},
			expected: struct {
				offset      int
				opCode      OpCode
				args        []byte
				description string
				stringRep   string
				protoRep    *opcode_pb.Instruction
			}{
				offset:      1,
				opCode:      STOP,
				args:        []byte{0x01, 0x02},
				description: "Test Description",
				stringRep:   "STOP", // Assuming STOP's string representation is "STOP"
				protoRep: &opcode_pb.Instruction{
					Offset:      1,
					OpCode:      opcode_pb.OpCode_STOP, // Assuming corresponding proto enum value
					Args:        []byte{0x01, 0x02},
					Description: "Test Description",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected.offset, tt.instruction.GetOffset())
			assert.Equal(t, tt.expected.opCode, tt.instruction.GetCode())
			assert.Equal(t, tt.expected.args, tt.instruction.GetArgs())
			assert.Equal(t, tt.expected.description, tt.instruction.GetDescription())
			assert.Equal(t, tt.expected.stringRep, tt.instruction.String())
			assert.Equal(t, tt.expected.protoRep, tt.instruction.ToProto())
		})
	}
}
