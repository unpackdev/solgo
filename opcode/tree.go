package opcode

import (
	"fmt"
	"strings"
)

// TreeNode represents a node in the opcode execution tree. Each node contains an instruction and
// potentially has multiple children, representing subsequent instructions.
type TreeNode struct {
	// Instruction represents the opcode instruction associated with this tree node.
	Instruction Instruction `json:"instruction"`
	// Children is a slice of TreeNode pointers, representing the child instructions of this node.
	Children []*TreeNode `json:"children"`
}

// GetTree constructs and returns the root of the opcode execution tree based on the instructions
// present in the Decompiler. If there are no instructions, it returns nil.
func (d *Decompiler) GetTree() *TreeNode {
	if len(d.instructions) == 0 {
		return nil
	}

	root := &TreeNode{}
	stack := []*TreeNode{root}
	d.buildTree(stack)
	return root.Children[0]
}

// Print outputs the structured representation of the opcode execution tree starting from the calling TreeNode.
// The output is indented for better readability.
func (t *TreeNode) Print() {
	t.printExecutionTree(t, 0)
}

// printExecutionTree is a recursive helper function that outputs the structured representation of the opcode
// execution tree starting from the provided node. The output is indented based on the indent parameter.
func (t *TreeNode) printExecutionTree(node *TreeNode, indent int) {
	fmt.Printf("%sOffset: 0x%04x, OpCode: %s, Args: %x\n", strings.Repeat(" ", indent*2), node.Instruction.Offset, node.Instruction.OpCode.String(), node.Instruction.Args)
	for _, child := range node.Children {
		t.printExecutionTree(child, indent+1)
	}
}

// buildTree constructs the opcode execution tree based on the instructions present in the Decompiler.
// It uses a stack-based approach to keep track of parent-child relationships between instructions.
func (d *Decompiler) buildTree(stack []*TreeNode) {
	for _, instruction := range d.instructions {
		node := &TreeNode{Instruction: instruction}

		if len(stack) > 0 {
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, node)
		}

		stack = append(stack, node)

		if instruction.OpCode.IsJump() {
			stack = stack[:len(stack)-1]
		}
	}
}
