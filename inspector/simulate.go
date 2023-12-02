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

func (m *SimulateDetector) Enter(ctx context.Context) (DetectorFn, error) {
	/* 	uniswap, err := bindings.NewUniswap(ctx, m.GetBindingManager(), bindings.DefaultUniswapBindOptions())
	   	if err != nil {
	   		zap.L().Error("failed to create uniswap bindings", zap.Error(err))
	   		return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
	   	}

	   	uniswap.EstimateTaxesForToken(m.GetAddress()) */

	//sim := m.Inspector.GetSimulator()

	// First thing what we need to do is figure out if this is in fact a token or not...
	m.GetReport()

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *SimulateDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *SimulateDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *SimulateDetector) Results() any {
	return m.results
}
