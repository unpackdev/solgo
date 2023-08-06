package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Event struct {
	unit *ast.EventDefinition

	Id         int64           `json:"id"`
	NodeType   ast_pb.NodeType `json:"node_type"`
	Name       string          `json:"name"`
	Anonymous  bool            `json:"anonymous"`
	Parameters []*Parameter    `json:"parameters"`
}

func (e *Event) GetAST() *ast.EventDefinition {
	return e.unit
}

// GetId returns the unique identifier of the node.
func (e *Event) GetId() int64 {
	return e.Id
}

// GetNodeType returns the type of the node in the AST.
func (e *Event) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetName returns the name of the node.
func (e *Event) GetName() string {
	return e.Name
}

// GetParameters returns the parameters of the event.
func (e *Event) GetParameters() []*Parameter {
	return e.Parameters
}

// IsAnonymous returns whether the event is anonymous.
func (e *Event) IsAnonymous() bool {
	return e.Anonymous
}

// GetSrc returns the source location of the node.
func (e *Event) GetSrc() ast.SrcNode {
	return e.unit.GetSrc()
}

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

func (b *Builder) processEvent(unit *ast.EventDefinition) *Event {
	toReturn := &Event{
		unit:       unit,
		Id:         unit.GetId(),
		NodeType:   unit.GetType(),
		Name:       unit.GetName(),
		Anonymous:  unit.IsAnonymous(),
		Parameters: make([]*Parameter, 0),
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		toReturn.Parameters = append(toReturn.Parameters, &Parameter{
			unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		})
	}

	return toReturn
}
