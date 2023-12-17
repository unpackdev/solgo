package inspector

import (
	"context"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type ExternalCallsResults struct {
	Detected          bool `json:"detected"`
	RenounceOwnership bool `json:"renounce_ownership"`
	CanLookupOwner    bool `json:"can_lookup_owner"`
	OwnerModifiable   bool `json:"owner_modifiable"`
	CanSelfDestruct   bool `json:"can_self_destruct"`
}

type ExternalCallsDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *ExternalCallsResults
}

func NewExternalCallsDetector(ctx context.Context, inspector *Inspector) Detector {
	return &ExternalCallsDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"call", "delegatecall", "staticcall", // Direct low-level calls
			"transfer", "transferFrom", // ERC20, ERC721 transfer functions
		},
		results: &ExternalCallsResults{},
	}
}

func (m *ExternalCallsDetector) Name() string {
	return "External Calls Detector"
}

func (m *ExternalCallsDetector) Type() DetectorType {
	return ExternalCallsDetectorType
}

// SetInspector sets the inspector for the detector
func (m *ExternalCallsDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *ExternalCallsDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *ExternalCallsDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_CALL: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.FunctionCall); ok {
				if expr, ok := fn.GetExpression().(*ast.PrimaryExpression); ok {
					if utils.StringInSlice(expr.GetName(), m.functionNames) {
						m.results.Detected = true
						return false, nil
					}
				}
			}
			return true, nil
		},
		ast_pb.NodeType_MEMBER_ACCESS: func(node ast.Node[ast.NodeType]) (bool, error) {
			if ma, ok := node.(*ast.MemberAccessExpression); ok {
				if expr, ok := ma.GetExpression().(*ast.PrimaryExpression); ok {
					if utils.StringInSlice(expr.GetName(), m.functionNames) {
						m.results.Detected = true
						return false, nil
					}

					if expr.GetTypeDescription() != nil && strings.Contains(expr.GetTypeDescription().GetString(), "contract") {
						m.results.Detected = true
						return false, nil
					}
				}

			}
			return true, nil
		},
	}, nil
}

func (m *ExternalCallsDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ExternalCallsDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ExternalCallsDetector) Results() any {
	return m.results
}
