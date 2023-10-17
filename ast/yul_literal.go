package ast

import (
	"encoding/hex"
	"fmt"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulLiteralStatement represents a Yul literal in the AST.
type YulLiteralStatement struct {
	*ASTBuilder                 // Embedded ASTBuilder for utility functions.
	Id          int64           `json:"id"`
	NodeType    ast_pb.NodeType `json:"node_type"`
	Kind        ast_pb.NodeType `json:"kind"`
	Src         SrcNode         `json:"src"`
	Value       string          `json:"value"`
	HexValue    string          `json:"hex_value"`
}

// NewYulLiteralStatement initializes a new YulLiteralStatement node.
func NewYulLiteralStatement(b *ASTBuilder) *YulLiteralStatement {
	return &YulLiteralStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_LITERAL,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulLiteralStatement node.
// Currently, this method always returns false and does not set any reference descriptor.
func (y *YulLiteralStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the ID of the YulLiteralStatement.
func (y *YulLiteralStatement) GetId() int64 {
	return y.Id
}

// GetType retrieves the node type of the YulLiteralStatement.
func (y *YulLiteralStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulLiteralStatement) GetKind() ast_pb.NodeType {
	return y.Kind
}

// GetSrc retrieves the source node information of the YulLiteralStatement.
func (y *YulLiteralStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes retrieves child nodes of the YulLiteralStatement.
// This returns an empty slice as YulLiteralStatement doesn't have child nodes.
func (y *YulLiteralStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	return toReturn
}

// GetTypeDescription retrieves the type description of the YulLiteralStatement.
// This currently returns an empty TypeDescription.
func (y *YulLiteralStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulLiteralStatement) GetValue() string {
	return y.Value
}

func (y *YulLiteralStatement) GetHexValue() string {
	return y.HexValue
}

// ToProto converts the YulLiteralStatement to its Protocol Buffer representation.
func (y *YulLiteralStatement) ToProto() NodeType {
	toReturn := ast_pb.YulLiteralStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
		Kind:     y.GetKind(),
		Value:    y.GetValue(),
		HexValue: y.GetHexValue(),
	}

	return NewTypedStruct(&toReturn, "YulLiteralStatement")
}

// Parse processes the given YulLiteralContext to populate the fields of the YulLiteralStatement.
func (y *YulLiteralStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulLiteralContext,
) Node[NodeType] {
	// Handle Boolean literals
	if ctx.YulBoolean() != nil {
		literal := ctx.YulBoolean()
		y.Value = literal.GetText()
		y.Kind = ast_pb.NodeType_BOOLEAN
		y.Src = SrcNode{
			Id:          y.GetNextID(),
			Line:        int64(literal.GetStart().GetLine()),
			Column:      int64(literal.GetStart().GetColumn()),
			Start:       int64(literal.GetStart().GetStart()),
			End:         int64(literal.GetStart().GetStop()),
			Length:      int64(literal.GetStart().GetStop() - literal.GetStart().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}
	}

	// Handle Decimal literals
	if ctx.YulDecimalNumber() != nil {
		literal := ctx.YulDecimalNumber()
		y.Value = literal.GetText()
		y.Kind = ast_pb.NodeType_DECIMAL_NUMBER
		y.Src = SrcNode{
			Id:          y.GetNextID(),
			Line:        int64(literal.GetSymbol().GetLine()),
			Column:      int64(literal.GetSymbol().GetColumn()),
			Start:       int64(literal.GetSymbol().GetStart()),
			End:         int64(literal.GetSymbol().GetStop()),
			Length:      int64(literal.GetSymbol().GetStop() - literal.GetSymbol().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}
	}

	// Handle String literals
	if ctx.YulStringLiteral() != nil {
		literal := ctx.YulStringLiteral()
		y.Value = literal.GetText()
		y.Kind = ast_pb.NodeType_STRING
		y.Src = SrcNode{
			Id:          y.GetNextID(),
			Line:        int64(literal.GetSymbol().GetLine()),
			Column:      int64(literal.GetSymbol().GetColumn()),
			Start:       int64(literal.GetSymbol().GetStart()),
			End:         int64(literal.GetSymbol().GetStop()),
			Length:      int64(literal.GetSymbol().GetStop() - literal.GetSymbol().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}
	}

	// Handle HexNumber literals
	if ctx.YulHexNumber() != nil {
		literal := ctx.YulHexNumber()
		y.Kind = ast_pb.NodeType_HEX_NUMBER
		y.HexValue = literal.GetText()
		y.Src = SrcNode{
			Id:          y.GetNextID(),
			Line:        int64(literal.GetSymbol().GetLine()),
			Column:      int64(literal.GetSymbol().GetColumn()),
			Start:       int64(literal.GetSymbol().GetStart()),
			End:         int64(literal.GetSymbol().GetStop()),
			Length:      int64(literal.GetSymbol().GetStop() - literal.GetSymbol().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}

		bytes, _ := hex.DecodeString(strings.Replace(y.HexValue, "0x", "", -1))
		value := int64(0)
		for _, b := range bytes {
			value = (value << 8) | int64(b)
		}
		y.Value = fmt.Sprintf("%d", value)
	}

	// Handle HexString literals
	if ctx.YulHexStringLiteral() != nil {
		literal := ctx.YulHexStringLiteral()
		y.Kind = ast_pb.NodeType_HEX_STRING
		y.HexValue = literal.GetText()
		y.Value = strings.Replace(y.HexValue, "0x", "", -1)
		y.Src = SrcNode{
			Id:          y.GetNextID(),
			Line:        int64(literal.GetSymbol().GetLine()),
			Column:      int64(literal.GetSymbol().GetColumn()),
			Start:       int64(literal.GetSymbol().GetStart()),
			End:         int64(literal.GetSymbol().GetStop()),
			Length:      int64(literal.GetSymbol().GetStop() - literal.GetSymbol().GetStart() + 1),
			ParentIndex: parentNode.GetId(),
		}
	}

	return y
}
