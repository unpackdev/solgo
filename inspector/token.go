package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type TokenResults struct {
	Detected bool `json:"detected"`
}

type TokenDetector struct {
	ctx context.Context
	*Inspector
	results *TokenResults
}

func NewTokenDetector(ctx context.Context, inspector *Inspector) Detector {
	return &TokenDetector{
		ctx:       ctx,
		Inspector: inspector,
		results:   &TokenResults{},
	}
}

func (m *TokenDetector) Name() string {
	return "State Variable Detector"
}

func (m *TokenDetector) Type() DetectorType {
	return TokenDetectorType
}

func (m *TokenDetector) Enter(ctx context.Context) (DetectorFn, error) {
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

func (m *TokenDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TokenDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TokenDetector) Results() any {
	return m.results
}
