package ast

import (
	"strings"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// The TupleExpression struct represents a tuple expression in Solidity.
type TupleExpression struct {
	// Embedding the ASTBuilder to provide common functionality
	*ASTBuilder

	// The unique identifier for the tuple expression
	Id int64 `json:"id"`

	// The type of the node, which is 'TUPLE_EXPRESSION' for a tuple expression
	NodeType ast_pb.NodeType `json:"node_type"`

	// The source information about the tuple expression, such as its line and column numbers in the source file
	Src SrcNode `json:"src"`

	// Whether the tuple expression is constant
	Constant bool `json:"is_constant"`

	// Whether the tuple expression is pure
	Pure bool `json:"is_pure"`

	// The components of the tuple expression
	Components []Node[NodeType] `json:"components"`

	// The referenced declaration of the tuple expression
	ReferencedDeclaration int64 `json:"referenced_declaration,omitempty"`

	// The type description of the tuple expression
	TypeDescription *TypeDescription `json:"type_description"`
}

// NewTupleExpression creates a new TupleExpression instance.
func NewTupleExpression(b *ASTBuilder) *TupleExpression {
	return &TupleExpression{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_TUPLE_EXPRESSION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the TupleExpression node.
func (t *TupleExpression) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	t.ReferencedDeclaration = refId
	t.TypeDescription = refDesc
	return false
}

// GetId returns the unique identifier of the tuple expression.
func (t *TupleExpression) GetId() int64 {
	return t.Id
}

// GetType returns the type of the node, which is 'TUPLE_EXPRESSION' for a tuple expression.
func (t *TupleExpression) GetType() ast_pb.NodeType {
	return t.NodeType
}

// GetSrc returns the source information about the tuple expression.
func (t *TupleExpression) GetSrc() SrcNode {
	return t.Src
}

// GetComponents returns the components of the tuple expression.
func (t *TupleExpression) GetComponents() []Node[NodeType] {
	return t.Components
}

// GetNodes returns the components of the tuple expression.
func (t *TupleExpression) GetNodes() []Node[NodeType] {
	return t.Components
}

// GetTypeDescription returns the type description of the tuple expression.
func (t *TupleExpression) GetTypeDescription() *TypeDescription {
	return t.TypeDescription
}

// IsConstant returns whether the tuple expression is constant.
func (t *TupleExpression) IsConstant() bool {
	return t.Constant
}

// IsPure returns whether the tuple expression is pure.
func (t *TupleExpression) IsPure() bool {
	return t.Pure
}

// GetReferencedDeclaration returns the referenced declaration of the tuple expression.
func (t *TupleExpression) GetReferencedDeclaration() int64 {
	return t.ReferencedDeclaration
}

// ToProto returns the protobuf representation of the tuple expression.
func (t *TupleExpression) ToProto() NodeType {
	proto := ast_pb.Tuple{
		Id:                    t.GetId(),
		NodeType:              t.GetType(),
		Src:                   t.GetSrc().ToProto(),
		IsConstant:            t.IsConstant(),
		IsPure:                t.IsPure(),
		ReferencedDeclaration: t.GetReferencedDeclaration(),
		TypeDescription:       t.GetTypeDescription().ToProto(),
	}

	for _, component := range t.GetComponents() {
		proto.Components = append(proto.Components, component.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Tuple")
}

// Parse parses a tuple expression from the provided parser.TupleContext and returns the corresponding TupleExpression.
func (t *TupleExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx *parser.TupleContext,
) Node[NodeType] {
	t.Src = SrcNode{
		Id:     t.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if exprNode != nil {
				return exprNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	expression := NewExpression(t.ASTBuilder)
	for _, tupleCtx := range ctx.TupleExpression().AllExpression() {
		expr := expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, t, tupleCtx)
		t.Components = append(
			t.Components,
			expr,
		)
		// A bit of a hack as we have interfaces but it works...
		switch exprCtx := expr.(type) {
		case *PrimaryExpression:
			if exprCtx.IsPure() {
				t.Pure = true
				break
			}
		}
	}

	t.TypeDescription = t.buildTypeDescription()

	t.dumpNode(t)
	return t
}

func (t *TupleExpression) buildTypeDescription() *TypeDescription {
	typeString := "tuple("
	typeIdentifier := "t_tuple_"
	typeStrings := make([]string, 0)
	typeIdentifiers := make([]string, 0)

	for _, component := range t.GetComponents() {
		td := component.GetTypeDescription()
		typeStrings = append(typeStrings, td.TypeString)
		typeIdentifiers = append(typeIdentifiers, "$_"+td.TypeIdentifier)
	}
	typeString += strings.Join(typeStrings, ",") + ")"
	typeIdentifier += strings.Join(typeIdentifiers, "_")
	typeIdentifier += "$"

	return &TypeDescription{
		TypeString:     typeString,
		TypeIdentifier: typeIdentifier,
	}
}
