package ast

import (
	"fmt"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type EventDefinition struct {
	*ASTBuilder
	SourceUnitName string          `json:"-"`
	Id             int64           `json:"id"`
	NodeType       ast_pb.NodeType `json:"node_type"`
	Src            SrcNode         `json:"src"`
	Parameters     *ParameterList  `json:"parameters"`
	Name           string          `json:"name"`
	Anonymous      bool            `json:"anonymous"`
}

func NewEventDefinition(b *ASTBuilder) *EventDefinition {
	return &EventDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_EVENT_DEFINITION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the EventDefinition node.
// We don't need to do any reference description updates here, at least for now...
func (e *EventDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (e *EventDefinition) GetId() int64 {
	return e.Id
}

func (e *EventDefinition) GetType() ast_pb.NodeType {
	return e.NodeType
}

func (e *EventDefinition) GetSrc() SrcNode {
	return e.Src
}

func (e *EventDefinition) GetName() string {
	return e.Name
}

func (e *EventDefinition) IsAnonymous() bool {
	return e.Anonymous
}

func (e *EventDefinition) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeIdentifier: fmt.Sprintf(
			"t_event&_%s_%s_&%d", e.SourceUnitName, e.GetName(), e.GetId(),
		),
		TypeString: fmt.Sprintf(
			"event %s.%s", e.SourceUnitName, e.GetName(),
		),
	}
}

func (e *EventDefinition) GetParameters() *ParameterList {
	return e.Parameters
}

func (e *EventDefinition) GetNodes() []Node[NodeType] {
	return nil
}

func (e *EventDefinition) ToProto() NodeType {
	proto := ast_pb.Event{
		Id:              e.GetId(),
		Name:            e.GetName(),
		NodeType:        e.GetType(),
		Src:             e.GetSrc().ToProto(),
		Anonymous:       e.IsAnonymous(),
		Parameters:      e.GetParameters().ToProto(),
		TypeDescription: e.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "Event")
}

func (e *EventDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.EventDefinitionContext,
) Node[NodeType] {
	e.SourceUnitName = unit.GetName()
	e.Src = SrcNode{
		Id:          e.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	e.Anonymous = ctx.Anonymous() != nil
	e.Name = ctx.Identifier().GetText()

	parameters := NewParameterList(e.ASTBuilder)
	parameters.ParseEventParameters(unit, e, ctx.AllEventParameter())
	e.Parameters = parameters

	e.currentEvents = append(e.currentEvents, e)
	return e
}
