package ast

type SrcNode struct {
	// Id is the unique identifier of the source node.
	Id int64 `json:"id"`
	// Line is the line of the source node.
	Line int64 `json:"line"`
	// Column is the column of the source node.
	Column int64 `json:"column"`
	// Start is the start of the source node.
	Start int64 `json:"start"`
	// End is the end of the source node.
	End int64 `json:"end"`
	// Length is the length of the source node.
	Length int64 `json:"length"`
	// ParentIndex is the index of the parent node.
	ParentIndex int64 `json:"parent_index"`
}

func NewSrcNode(id int64, line int64, column int64, start int64, end int64, length int64, parentIndex int64) SrcNode {
	return SrcNode{
		Id:          id,
		Line:        line,
		Column:      column,
		Start:       start,
		End:         end,
		Length:      length,
		ParentIndex: parentIndex,
	}
}

// GetId returns the ID of the source node.
func (s SrcNode) GetId() int64 {
	return s.Id
}

// GetLine returns the line of the source node.
func (s SrcNode) GetLine() int64 {
	return s.Line
}

// GetColumn returns the column of the source node.
func (s SrcNode) GetColumn() int64 {
	return s.Column
}

// GetStart returns the start of the source node.
func (s SrcNode) GetStart() int64 {
	return s.Start
}

// GetEnd returns the end of the source node.
func (s SrcNode) GetEnd() int64 {
	return s.End
}

// GetLength returns the length of the source node.
func (s SrcNode) GetLength() int64 {
	return s.Length
}

// GetParentIndex returns the index of the parent node.
func (s SrcNode) GetParentIndex() int64 {
	return s.ParentIndex
}
