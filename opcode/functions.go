package opcode

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// FunctionTreeNode represents a node in the opcode execution tree that represents a function.
type FunctionTreeNode struct {
	FunctionSignatureHex string `json:"functionSignatureHex"`
	FunctionSignature    string `json:"function_signature"`
	FunctionBytesHex     string `json:"function_bytes_hex"`
	FunctionBytes        []byte `json:"function_bytes"`
	HasFunctionSignature bool   `json:"has_function_signature"`
	*TreeNode
}

// GetFunctions extracts all functions from the opcode execution tree and returns them as a slice of FunctionTreeNode pointers.
func (d *Decompiler) GetFunctions() []*FunctionTreeNode {
	if len(d.instructions) == 0 {
		return nil
	}

	var functions []*FunctionTreeNode
	stack := []*FunctionTreeNode{nil}
	d.buildFunctions(stack, &functions)
	return functions
}

// buildFunctions traverses the opcode execution tree to identify and extract functions.
func (d *Decompiler) buildFunctions(stack []*FunctionTreeNode, functions *[]*FunctionTreeNode) {
	for _, instruction := range d.instructions {
		if instruction.OpCode.IsFunctionStart() {
			functionSignatureBytes := d.calculateFunctionSignature(instruction.Offset)

			functionSignature := ""
			if len(functionSignatureBytes) >= 4 {
				functionSignature = common.Bytes2Hex(functionSignatureBytes[:4])
			}

			functionNode := &FunctionTreeNode{
				FunctionSignature:    functionSignature,
				FunctionSignatureHex: fmt.Sprintf("0x%s", functionSignature),
				FunctionBytesHex:     common.Bytes2Hex(functionSignatureBytes),
				FunctionBytes:        functionSignatureBytes,
				HasFunctionSignature: len(functionSignatureBytes) > 0,
				TreeNode:             &TreeNode{Instruction: instruction, Children: make([]*TreeNode, 0)},
			}

			*functions = append(*functions, functionNode)
			stack = append(stack, functionNode)
		} else if instruction.OpCode.IsFunctionEnd() {
			if len(stack) == 0 {
				fmt.Println("Error: Found function end without corresponding start.")
				continue
			}
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) == 0 {
				fmt.Println("Error: Found instruction outside of a function.")
				continue
			}

			if parent := stack[len(stack)-1]; parent != nil {
				parent.Children = append(parent.Children, &TreeNode{Instruction: instruction, Children: make([]*TreeNode, 0)})
			}
		}
	}
}

// calculateFunctionSignature calculates the function signature from the EVM bytecode at the given offset.
func (d *Decompiler) calculateFunctionSignature(offset int) []byte {
	var targetInstruction *Instruction
	for i := offset - 1; i >= 0; i-- {
		if i >= len(d.instructions) {
			break
		}
		instr := d.instructions[i]
		if instr.OpCode == JUMPDEST {
			targetInstruction = &instr
			break
		}
	}

	if targetInstruction == nil {
		return nil
	}

	var signature []byte
	for i := offset; i < len(d.instructions); i++ {
		instr := d.instructions[i]
		if instr.OpCode.IsPush() && len(instr.Args) == 4 {
			signature = instr.Args
			break
		}
	}

	return signature
}
