package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type PausableResults struct {
	Detected          bool `json:"detected"`
	RenounceOwnership bool `json:"renounce_ownership"`
	CanLookupOwner    bool `json:"can_lookup_owner"`
	OwnerModifiable   bool `json:"owner_modifiable"`
	CanSelfDestruct   bool `json:"can_self_destruct"`
}

type PausableDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *PausableResults
}

func NewPausableDetector(ctx context.Context, inspector *Inspector) Detector {
	return &PausableDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"paused", "pause", "unpause", "_pause", "_unpause",
			"isPaused",
			"canPause", "canUnpause",
			"togglePause",
			"setPaused",
			"enablePause", "disablePause",
			"pauseTransfer", "unpauseTransfer",
			"pauseAll", "unpauseAll",
		},
		results: &PausableResults{},
	}
}

func (m *PausableDetector) Name() string {
	return "Pausable Detector"
}

func (m *PausableDetector) Type() DetectorType {
	return PausableDetectorType
}

// SetInspector sets the inspector for the detector
func (m *PausableDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *PausableDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *PausableDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fn.GetName(), m.functionNames) {
					m.results.Detected = true
					return false, nil
				}
			}
			return true, nil
		},
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
	}, nil
}

func (m *PausableDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *PausableDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *PausableDetector) Results() any {
	return m.results
}
