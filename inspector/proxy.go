package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/utils"
)

type ProxyResults struct {
	Detected bool         `json:"detected"`
	Standard *ir.Standard `json:"standard"`
}

type ProxyDetector struct {
	ctx           context.Context
	inspector     *Inspector
	enabled       bool
	functionNames []string
	results       *ProxyResults
}

func NewProxyDetector(ctx context.Context, inspector *Inspector) Detector {
	return &ProxyDetector{
		ctx:       ctx,
		inspector: inspector,
		enabled:   false,
		functionNames: []string{
			"upgradeTo", "upgradeToAndCall", "getAdmin", "changeAdmin", "_implementation", "implementation",
			"upgradeBeaconToAndCall", "getImplementation", "_setImplementation", "_setAdmin", "_dispatchUpgradeToAndCall",
			"_delegate",
		},
		results: &ProxyResults{},
	}
}

func (m *ProxyDetector) Name() string {
	return "Proxy Detector"
}

func (m *ProxyDetector) Type() DetectorType {
	return ProxyDetectorType
}

// SetInspector sets the inspector for the detector
func (m *ProxyDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *ProxyDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *ProxyDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ProxyDetector) Detect(ctx context.Context) (DetectorFn, error) {
	// This detector can use IR as well to do its job and walking through the AST if absolutely necessary...

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fn.GetName(), m.functionNames) {
					m.results.Detected = true
				}
			}
			return true, nil
		},
	}, nil
}

func (m *ProxyDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ProxyDetector) Results() any {
	return m.results
}
