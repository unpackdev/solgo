package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type VariableDeclaration struct {
	*ASTBuilder

	Id           int64           `json:"id"`
	NodeType     ast_pb.NodeType `json:"node_type"`
	Src          SrcNode         `json:"src"`
	Assignments  []int64         `json:"assignments"`
	Declarations []*Declaration  `json:"declarations"`
	InitialValue Node[NodeType]  `json:"initial_value,omitempty"`
}

func NewVariableDeclarationStatement(b *ASTBuilder) *VariableDeclaration {
	return &VariableDeclaration{
		ASTBuilder:   b,
		Id:           b.GetNextID(),
		NodeType:     ast_pb.NodeType_VARIABLE_DECLARATION,
		Assignments:  make([]int64, 0),
		Declarations: make([]*Declaration, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the VariableDeclaration node.
func (v *VariableDeclaration) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (v *VariableDeclaration) GetId() int64 {
	return v.Id
}

func (v *VariableDeclaration) GetType() ast_pb.NodeType {
	return v.NodeType
}

func (v *VariableDeclaration) GetSrc() SrcNode {
	return v.Src
}

func (v *VariableDeclaration) GetAssignments() []int64 {
	return v.Assignments
}

func (v *VariableDeclaration) GetDeclarations() []*Declaration {
	return v.Declarations
}

func (v *VariableDeclaration) GetInitialValue() Node[NodeType] {
	return v.InitialValue
}

func (v *VariableDeclaration) GetTypeDescription() *TypeDescription {
	if len(v.Declarations) > 0 {
		return v.Declarations[0].GetTypeDescription()
	}

	return nil
}

func (v *VariableDeclaration) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{v.InitialValue}
	for _, declaration := range v.GetDeclarations() {
		toReturn = append(toReturn, declaration)
	}

	return toReturn
}

func (v *VariableDeclaration) ToProto() NodeType {
	proto := ast_pb.Variable{
		Id:           v.Id,
		NodeType:     v.NodeType,
		Src:          v.Src.ToProto(),
		Assignments:  v.Assignments,
		Declarations: make([]*ast_pb.Declaration, 0),
	}

	if v.GetInitialValue() != nil {
		proto.InitialValue = v.GetInitialValue().ToProto().(*v3.TypedStruct)
	}

	for _, declaration := range v.GetDeclarations() {
		proto.Declarations = append(
			proto.Declarations,
			declaration.ToProto().(*ast_pb.Declaration),
		)
	}

	return NewTypedStruct(&proto, "Variable")
}

func (v *VariableDeclaration) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.VariableDeclarationStatementContext,
) {
	v.Src = SrcNode{
		Id:          v.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	if ctx.VariableDeclaration() != nil {
		declaration := NewDeclaration(v.ASTBuilder)
		declaration.ParseVariableDeclaration(unit, contractNode, fnNode, bodyNode, v, ctx.VariableDeclaration())
		v.Declarations = append(v.Declarations, declaration)
		v.Assignments = append(v.Assignments, declaration.GetId())
	}

	if ctx.Expression() != nil {
		expression := NewExpression(v.ASTBuilder)
		v.InitialValue = expression.Parse(unit, contractNode, fnNode, bodyNode, v, nil, ctx.Expression())
	}

	v.currentVariables = append(v.currentVariables, v)
}
