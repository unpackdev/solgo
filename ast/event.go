package ast

import (
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// EventDefinition represents an event definition in the Solidity abstract syntax tree (AST).
type EventDefinition struct {
	*ASTBuilder                    // Embedding the ASTBuilder for common functionality
	SourceUnitName string          `json:"-"`
	Id             int64           `json:"id"`         // Unique identifier for the event definition
	NodeType       ast_pb.NodeType `json:"node_type"`  // Type of the node (EVENT_DEFINITION for event definition)
	Src            SrcNode         `json:"src"`        // Source information about the event definition
	Parameters     *ParameterList  `json:"parameters"` // Parameters of the event
	Name           string          `json:"name"`       // Name of the event
	Anonymous      bool            `json:"anonymous"`  // Indicates if the event is anonymous
}

// NewEventDefinition creates a new EventDefinition instance.
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

// GetId returns the unique identifier of the event definition.
func (e *EventDefinition) GetId() int64 {
	return e.Id
}

// GetType returns the type of the node, which is 'EVENT_DEFINITION' for an event definition.
func (e *EventDefinition) GetType() ast_pb.NodeType {
	return e.NodeType
}

// GetSrc returns the source information about the event definition.
func (e *EventDefinition) GetSrc() SrcNode {
	return e.Src
}

// GetName returns the name of the event.
func (e *EventDefinition) GetName() string {
	return e.Name
}

// IsAnonymous returns whether the event is anonymous.
func (e *EventDefinition) IsAnonymous() bool {
	return e.Anonymous
}

// GetTypeDescription returns the type description of the event.
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

// GetParameters returns the parameters of the event.
func (e *EventDefinition) GetParameters() *ParameterList {
	return e.Parameters
}

// GetNodes returns the nodes representing the parameters of the event.
func (e *EventDefinition) GetNodes() []Node[NodeType] {
	return e.Parameters.GetNodes()
}

// ToProto returns the protobuf representation of the event definition.
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

// Parse parses an event definition from the provided parser.EventDefinitionContext and updates the current instance.
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
