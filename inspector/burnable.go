package inspector

import (
	"context"
	"fmt"

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
	return "Burnable Detector"
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

func (m *BurnDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

func (m *BurnDetector) GetInspector() *Inspector {
	return m.inspector
}

// Enter for now does nothing for mint detector. It may be needed in the future.
func (m *BurnDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *BurnDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
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
			return true, nil
		},
	}, nil
}

func (m *BurnDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){

		// Problem is that mint function can be discovered at any point in time so we need to go one more time
		// through whole process in case that mint is discovered in order to get all of the reference locations where
		// mint function is being called out...
		// Burn function can exist and never be used or it can be announced as not being used where we in fact see that
		// it can be used....
		ast_pb.NodeType_FUNCTION_DEFINITION: func(fnNode ast.Node[ast.NodeType]) (bool, error) {
			// We do not want to continue if we did not discover mint function...
			if !m.results.Detected {
				return false, nil
			}

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_MEMBER_ACCESS, func(node ast.Node[ast.NodeType]) (bool, error) {
				nodeCtx, ok := node.(*ast.MemberAccessExpression)
				if !ok {
					return true, fmt.Errorf("unable to convert node to MemberAccessExpression type in BurnDetector.Exit: %T", node)
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

				return true, nil
			})

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_FUNCTION_CALL, func(node ast.Node[ast.NodeType]) (bool, error) {
				nodeCtx, ok := node.(*ast.FunctionCall)
				if !ok {
					return true, fmt.Errorf("unable to convert node to FunctionCall type in BurnDetector.Exit: %T", node)
				}

				expressionCtx, ok := nodeCtx.GetExpression().(*ast.PrimaryExpression)
				if !ok {
					return true, fmt.Errorf("unable to convert node to PrimaryExpression type in BurnDetector.Exit: %T", nodeCtx.GetExpression())
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
				return true, nil
			})

			return true, nil
		},
	}, nil
}

func (m *BurnDetector) Results() any {
	return m.results
}
