package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type VariableDeclaration struct {
	Name          string                 `json:"name"`
	StateVariable bool                   `json:"state_variable"`
	Constant      bool                   `json:"constant"`
	Statement     ast.Node[ast.NodeType] `json:"statement"`
}

type StateVariableResults struct {
	Detected     bool                   `json:"detected"`
	Declarations []*VariableDeclaration `json:"declarations"`
}

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
		results: &StateVariableResults{
			Declarations: make([]*VariableDeclaration, 0),
		},
	}
}

func (m *StateVariableDetector) Name() string {
	return "StateVariable Detector"
}

func (m *StateVariableDetector) Type() DetectorType {
	return StateVariableDetectorType
}

func (m *StateVariableDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
		ast_pb.NodeType_VARIABLE_DECLARATION: func(node ast.Node[ast.NodeType]) bool {
			if varCtx, ok := node.(*ast.VariableDeclaration); ok {
				for _, declaration := range varCtx.GetDeclarations() {
					m.results.Declarations = append(m.results.Declarations, &VariableDeclaration{
						Name:          declaration.GetName(),
						StateVariable: declaration.GetIsStateVariable(),
						//Statement:     declaration,
					})
					//m.results.StateVariable = varCtx
				}
			} else if varCtx, ok := node.(*ast.StateVariableDeclaration); ok {
				m.results.Declarations = append(m.results.Declarations, &VariableDeclaration{
					Name:          varCtx.GetName(),
					StateVariable: varCtx.IsStateVariable(),
					Constant:      varCtx.IsConstant(),
					Statement:     varCtx,
				})
			}

			return true
		},
	}
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
