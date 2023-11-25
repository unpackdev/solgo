package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type StateVariableResults struct{}

type StateVariableDetector struct {
	ctx       context.Context
	inspector *Inspector
	enabled   bool
	results   *StateVariableResults
}

func NewStateVariableDetector(ctx context.Context, inspector *Inspector) Detector {
	return &StateVariableDetector{
		ctx:       ctx,
		inspector: inspector,
		enabled:   false,
		results:   &StateVariableResults{},
	}
}

func (m *StateVariableDetector) Name() string {
	return "StateVariable Detector"
}

func (m *StateVariableDetector) Type() DetectorType {
	return StateVariableDetectorType
}

func (m *StateVariableDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *StateVariableDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *StateVariableDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *StateVariableDetector) Results() any {
	return m.results
}
