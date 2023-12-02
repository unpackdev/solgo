package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/standards"
)

type ProxyResults struct {
	Detected bool         `json:"detected"`
	Standard *ir.Standard `json:"standard"`
}

type ProxyDetector struct {
	ctx       context.Context
	inspector *Inspector
	enabled   bool
	results   *ProxyResults
}

func NewProxyDetector(ctx context.Context, inspector *Inspector) Detector {
	return &ProxyDetector{
		ctx:       ctx,
		inspector: inspector,
		enabled:   false,
		results:   &ProxyResults{},
	}
}

func (m *ProxyDetector) Name() string {
	return "Proxy Detector"
}

func (m *ProxyDetector) Type() DetectorType {
	return ProxyDetectorType
}

func (m *ProxyDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ProxyDetector) Detect(ctx context.Context) (DetectorFn, error) {
	// This detector can use IR as well to do its job and walking through the AST if absolutely necessary...
	detector := m.inspector.GetDetector()
	irRoot := detector.GetIR().GetRoot()

	if irRoot.HasHighConfidenceStandard(standards.ERC1967) {
		m.results.Detected = true
		m.results.Standard = irRoot.GetStandard(standards.ERC1967)
	} else if irRoot.HasHighConfidenceStandard(standards.ERC1822) {
		m.results.Detected = true
		m.results.Standard = irRoot.GetStandard(standards.ERC1822)
	}

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ProxyDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *ProxyDetector) Results() any {
	return m.results
}
