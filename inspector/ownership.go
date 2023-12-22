package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type OwnershipResults struct {
	Detected          bool `json:"detected"`
	RenounceOwnership bool `json:"renounce_ownership"`
	CanLookupOwner    bool `json:"can_lookup_owner"`
	OwnerModifiable   bool `json:"owner_modifiable"`
	CanSelfDestruct   bool `json:"can_self_destruct"`
}

type OwnershipDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *OwnershipResults
}

func NewOwnershipDetector(ctx context.Context, inspector *Inspector) Detector {
	return &OwnershipDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"transferOwnership", "renounceOwnership", "_transferOwnership", "_renounceOwnership",
			"owner", "setOwner", "claimOwnership", "initializeOwnership", "selfdestruct", "setTokenOwner",
			"confirmOwnershipTransfer", "cancelOwnershipTransfer",
		},
		results: &OwnershipResults{},
	}
}

func (m *OwnershipDetector) Name() string {
	return "State Variable Detector"
}

func (m *OwnershipDetector) Type() DetectorType {
	return OwnershipDetectorType
}

// SetInspector sets the inspector for the detector
func (m *OwnershipDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *OwnershipDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *OwnershipDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fn.GetName(), m.functionNames) {
					m.results.Detected = true

					if fn.GetName() == "setOwner" || fn.GetName() == "transferOwnership" || fn.GetName() == "_transferOwnership" || fn.GetName() == "claimOwnership" || fn.GetName() == "setTokenOwner" {
						m.results.OwnerModifiable = true
						return true, nil
					}

					if fn.GetName() == "renounceOwnership" || fn.GetName() == "_renounceOwnership" || fn.GetName() == "RenounceOwner" {
						m.results.RenounceOwnership = true
						return true, nil
					}

					if fn.GetName() == "owner" {
						m.results.CanLookupOwner = true
						return true, nil
					}
				}
			}
			return true, nil
		},
		ast_pb.NodeType_FUNCTION_CALL: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.FunctionCall); ok {
				if expr, ok := fn.GetExpression().(*ast.PrimaryExpression); ok {
					if expr.GetName() == "selfdestruct" {
						m.results.CanSelfDestruct = true
						return true, nil
					}
				}
			}
			return true, nil
		},
	}, nil
}

func (m *OwnershipDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *OwnershipDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *OwnershipDetector) Results() any {
	return m.results
}
