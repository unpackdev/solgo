package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type BurnResults struct {
	Detected                   bool                     `json:"detected"`
	FunctionName               string                   `json:"function_name"`
	Visibility                 ast_pb.Visibility        `json:"visibility"`
	ExternallyCallable         bool                     `json:"externally_callable"`
	ExternallCallableLocations []ast.Node[ast.NodeType] `json:"externally_callable_locations"`
	UsesConstructor            bool                     `json:"uses_constructor"`
	Constructor                *ast.Constructor         `json:"constructor"`
	Function                   *ast.Function            `json:"function"`
}

func (m BurnResults) IsDetected() bool {
	return m.Detected
}

func (m BurnResults) IsVisible() bool {
	return m.Visibility == ast_pb.Visibility_PUBLIC || m.Visibility == ast_pb.Visibility_EXTERNAL
}

type BurnDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *BurnResults
}

func NewBurnDetector(ctx context.Context, inspector *Inspector) Detector {
	return &BurnDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"burn", "burnFor", "burnTo", "burnBatch", "burnBatchFor", "burnBatchTo",
			"_burn", "_burnFor", "_burnTo", "_burnBatch", "_burnBatchFor", "_burnBatchTo",
		},
		results: &BurnResults{
			ExternallCallableLocations: make([]ast.Node[ast.NodeType], 0),
		},
	}
}

func (m *BurnDetector) Name() string {
	return "Burn Detector"
}

func (m *BurnDetector) Type() DetectorType {
	return BurnDetectorType
}

func (m *BurnDetector) RegisterFunctionName(fnName string) bool {
	if !utils.StringInSlice(fnName, m.functionNames) {
		m.functionNames = append(m.functionNames, fnName)
		return true
	}

	return false
}

func (m *BurnDetector) GetFunctionNames() []string {
	return m.functionNames
}

func (m *BurnDetector) FunctionNameExists(fnName string) bool {
	return utils.StringInSlice(fnName, m.functionNames)
}

// Enter for now does nothing for mint detector. It may be needed in the future.
func (m *BurnDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *BurnDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
			switch nodeCtx := node.(type) {
			case *ast.Constructor:
			case *ast.Function:
				if m.FunctionNameExists(nodeCtx.GetName()) {
					m.results.Detected = true
					m.results.Visibility = nodeCtx.GetVisibility()
					m.results.FunctionName = nodeCtx.GetName()
					//m.results.Function = nodeCtx
				}
			}
			return true
		},
	}
}

func (m *BurnDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{

		// Problem is that mint function can be discovered at any point in time so we need to go one more time
		// through whole process in case that mint is discovered in order to get all of the reference locations where
		// mint function is being called out...
		// Burn function can exist and never be used or it can be announced as not being used where we in fact see that
		// it can be used....
		ast_pb.NodeType_FUNCTION_DEFINITION: func(fnNode ast.Node[ast.NodeType]) bool {
			// We do not want to continue if we did not discover mint function...
			if !m.results.Detected {
				return false
			}

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_MEMBER_ACCESS, func(node ast.Node[ast.NodeType]) bool {
				nodeCtx, ok := node.(*ast.MemberAccessExpression)
				if !ok {
					return true
				}

				if nodeCtx.GetMemberName() == m.results.FunctionName {
					switch fnCtx := fnNode.(type) {
					case *ast.Function:
						if fnCtx.GetVisibility() == ast_pb.Visibility_PUBLIC || fnCtx.GetVisibility() == ast_pb.Visibility_EXTERNAL {
							m.results.ExternallyCallable = true
							m.results.ExternallCallableLocations = append(m.results.ExternallCallableLocations, fnNode)
						} else {
							// TODO: This should recursively look for other functions when internal or private function is visibility type
						}
					}
				}

				return true
			})

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_FUNCTION_CALL, func(node ast.Node[ast.NodeType]) bool {
				nodeCtx, ok := node.(*ast.FunctionCall)
				if !ok {
					return true
				}

				expressionCtx, ok := nodeCtx.GetExpression().(*ast.PrimaryExpression)
				if !ok {
					return true
				}

				if expressionCtx.GetName() == m.results.FunctionName {
					switch fnCtx := fnNode.(type) {
					case *ast.Constructor:
						m.results.UsesConstructor = true
						//m.results.Constructor = fnCtx
					case *ast.Function:
						if fnCtx.GetVisibility() == ast_pb.Visibility_PUBLIC || fnCtx.GetVisibility() == ast_pb.Visibility_EXTERNAL {
							m.results.ExternallyCallable = true
							m.results.ExternallCallableLocations = append(m.results.ExternallCallableLocations, fnNode)
						} else {
							// TODO: This should recursively look for other functions when internal or private function is visibility type
						}
					}
				}
				return true
			})

			return true
		},
	}
}

func (m *BurnDetector) Results() any {
	return m.results
}
