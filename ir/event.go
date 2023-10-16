package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Event represents an event definition in the IR.
type Event struct {
	Unit       *ast.EventDefinition `json:"ast"`
	Id         int64                `json:"id"`
	NodeType   ast_pb.NodeType      `json:"node_type"`
	Name       string               `json:"name"`
	Anonymous  bool                 `json:"anonymous"`
	Parameters []*Parameter         `json:"parameters"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the event definition.
func (e *Event) GetAST() *ast.EventDefinition {
	return e.Unit
}

// GetId returns the ID of the event definition.
func (e *Event) GetId() int64 {
	return e.Id
}

// GetNodeType returns the NodeType of the event definition.
func (e *Event) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetName returns the name of the event definition.
func (e *Event) GetName() string {
	return e.Name
}

// GetParameters returns the parameters of the event definition.
func (e *Event) GetParameters() []*Parameter {
	return e.Parameters
}

// IsAnonymous returns whether the event definition is anonymous.
func (e *Event) IsAnonymous() bool {
	return e.Anonymous
}

// GetSrc returns the source location of the event definition.
func (e *Event) GetSrc() ast.SrcNode {
	return e.Unit.GetSrc()
}

// ToProto converts the Event to its protobuf representation.
func (e *Event) ToProto() *ir_pb.Event {
	proto := &ir_pb.Event{
		Id:         e.GetId(),
		NodeType:   e.GetNodeType(),
		Name:       e.GetName(),
		Anonymous:  e.IsAnonymous(),
		Parameters: make([]*ir_pb.Parameter, 0),
	}

	for _, parameter := range e.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	return proto
}

// processEvent processes the event definition unit and returns the Event.
func (b *Builder) processEvent(unit *ast.EventDefinition) *Event {
	toReturn := &Event{
		Unit:       unit,
		Id:         unit.GetId(),
		NodeType:   unit.GetType(),
		Name:       unit.GetName(),
		Anonymous:  unit.IsAnonymous(),
		Parameters: make([]*Parameter, 0),
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		toReturn.Parameters = append(toReturn.Parameters, &Parameter{
			Unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Indexed:         parameter.IsIndexed(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	return toReturn
}
