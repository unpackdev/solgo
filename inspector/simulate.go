package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type SimulateResults struct {
	Detected bool `json:"detected"`
}

type SimulateDetector struct {
	ctx context.Context
	*Inspector
	results *SimulateResults
}

func NewSimulateDetector(ctx context.Context, inspector *Inspector) Detector {
	return &SimulateDetector{
		ctx:       ctx,
		Inspector: inspector,
		results:   &SimulateResults{},
	}
}

func (m *SimulateDetector) Name() string {
	return "State Variable Detector"
}

func (m *SimulateDetector) Type() DetectorType {
	return SimulateDetectorType
}

func (m *SimulateDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	/* 	uniswap, err := bindings.NewUniswap(ctx, m.GetBindingManager(), bindings.DefaultUniswapBindOptions())
	   	if err != nil {
	   		zap.L().Error("failed to create uniswap bindings", zap.Error(err))
	   		return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
	   	}

	   	uniswap.EstimateTaxesForToken(m.GetAddress()) */

	//sim := m.Inspector.GetSimulator()

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *SimulateDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *SimulateDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *SimulateDetector) Results() any {
	return m.results
}
