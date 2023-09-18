package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// SrcNode represents a node in the source code.
type SrcNode struct {
	Id          int64 `json:"id"`           // Unique identifier of the source node.
	Line        int64 `json:"line"`         // Line number of the source node in the source code.
	Column      int64 `json:"column"`       // Column number of the source node in the source code.
	Start       int64 `json:"start"`        // Start position of the source node in the source code.
	End         int64 `json:"end"`          // End position of the source node in the source code.
	Length      int64 `json:"length"`       // Length of the source node in the source code.
	ParentIndex int64 `json:"parent_index"` // Index of the parent node in the source code.
}

// GetId returns the unique identifier of the source node.
func (s SrcNode) GetId() int64 {
	return s.Id
}

// GetLine returns the line number of the source node in the source code.
func (s SrcNode) GetLine() int64 {
	return s.Line
}

// GetColumn returns the column number of the source node in the source code.
func (s SrcNode) GetColumn() int64 {
	return s.Column
}

// GetStart returns the start position of the source node in the source code.
func (s SrcNode) GetStart() int64 {
	return s.Start
}

// GetEnd returns the end position of the source node in the source code.
func (s SrcNode) GetEnd() int64 {
	return s.End
}

// GetLength returns the length of the source node in the source code.
func (s SrcNode) GetLength() int64 {
	return s.Length
}

// GetParentIndex returns the index of the parent node in the source code.
func (s SrcNode) GetParentIndex() int64 {
	return s.ParentIndex
}

// ToProto converts the SrcNode to a protocol buffer representation.
func (s SrcNode) ToProto() *ast_pb.Src {
	return &ast_pb.Src{
		Id:          s.GetId(),
		Line:        s.GetLine(),
		Column:      s.GetColumn(),
		Start:       s.GetStart(),
		End:         s.GetEnd(),
		Length:      s.GetLength(),
		ParentIndex: s.GetParentIndex(),
	}
}
