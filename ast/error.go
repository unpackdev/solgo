package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ErrorDefinition struct {
	*ASTBuilder
	SourceUnitName  string           `json:"-"`
	Id              int64            `json:"id"`
	NodeType        ast_pb.NodeType  `json:"node_type"`
	Src             SrcNode          `json:"src"`
	Name            string           `json:"name"`
	Parameters      *ParameterList   `json:"parameters"`
	TypeDescription *TypeDescription `json:"type_description"`
}

func NewErrorDefinition(b *ASTBuilder) *ErrorDefinition {
	return &ErrorDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_ERROR_DEFINITION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ErrorDefinition node.
// We don't need to do any reference description updates here, at least for now...
func (e *ErrorDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (e *ErrorDefinition) GetId() int64 {
	return e.Id
}

func (e *ErrorDefinition) GetType() ast_pb.NodeType {
	return e.NodeType
}

func (e *ErrorDefinition) GetSrc() SrcNode {
	return e.Src
}

func (e *ErrorDefinition) GetName() string {
	return e.Name
}

func (e *ErrorDefinition) GetTypeDescription() *TypeDescription {
	return e.TypeDescription
}

func (e *ErrorDefinition) GetParameters() *ParameterList {
	return e.Parameters
}

func (e *ErrorDefinition) GetSourceUnitName() string {
	return e.SourceUnitName
}

func (e *ErrorDefinition) GetNodes() []Node[NodeType] {
	return nil
}

func (e *ErrorDefinition) ToProto() NodeType {
	return ast_pb.Error{}
}

func (e *ErrorDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.ErrorDefinitionContext,
) Node[NodeType] {
	e.Src = SrcNode{
		Id:          e.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	e.SourceUnitName = unit.GetName()
	e.Name = ctx.GetName().GetText()

	e.TypeDescription = &TypeDescription{
		TypeIdentifier: fmt.Sprintf(
			"t_error$_%s_%s_$%d", e.SourceUnitName, e.GetName(), e.GetId(),
		),
		TypeString: fmt.Sprintf(
			"error %s.%s", e.SourceUnitName, e.GetName(),
		),
	}

	parameters := NewParameterList(e.ASTBuilder)
	parameters.ParseErrorParameters(unit, e, ctx.AllErrorParameter())
	e.Parameters = parameters

	e.currentErrors = append(e.currentErrors, e)
	return e
}
