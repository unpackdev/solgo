package inspector

import (
	"context"
	"math/big"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/standards"
	"go.uber.org/zap"
)

type TokenResults struct {
	*bindings.Token

	Detected    bool     `json:"detected"`
	Corrupted   bool     `json:"corrupted"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"total_supply"`
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
	return "Token (ERC-20) Detector"
}

func (m *TokenDetector) Type() DetectorType {
	return TokenDetectorType
}

func (m *TokenDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TokenDetector) Detect(ctx context.Context) (DetectorFn, error) {
	if m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil {
		//irRoot := m.GetDetector().GetIR().GetRoot()
		report := m.GetReport()

		if report.HasDetector(StandardsDetectorType) {
			if standardsDetector, ok := report.GetDetector(StandardsDetectorType).(*StandardsResults); ok {
				for _, standard := range standardsDetector.StandardTypes {
					if standard == standards.ERC20 {
						m.results.Detected = true

						tokenBinding, err := bindings.NewToken(ctx, report.Network, m.GetBindingManager(), bindings.DefaultTokenBindOptions(m.GetAddress()))
						if err != nil {
							zap.L().Error(
								"failed to create token bindings",
								zap.Error(err),
								zap.String("address", m.GetAddress().Hex()),
								zap.Any("network", report.Network),
							)
							return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
						}

						m.results.Token = tokenBinding

						name, err := tokenBinding.GetName()
						if err != nil {
							zap.L().Error(
								"failed to get token name",
								zap.Error(err),
								zap.String("address", m.GetAddress().Hex()),
								zap.Any("network", report.Network),
							)
						} else {
							m.results.Name = name
						}

						symbol, err := tokenBinding.GetSymbol()
						if err != nil {
							zap.L().Error(
								"failed to get token symbol",
								zap.Error(err),
								zap.String("address", m.GetAddress().Hex()),
								zap.Any("network", report.Network),
							)
							m.results.Corrupted = true
						} else {
							m.results.Symbol = symbol
						}

						decimals, err := tokenBinding.GetDecimals()
						if err != nil {
							zap.L().Error(
								"failed to get token decimals",
								zap.Error(err),
								zap.String("address", m.GetAddress().Hex()),
								zap.Any("network", report.Network),
							)
							m.results.Corrupted = true
						} else {
							m.results.Decimals = decimals
						}

						totalSupply, err := tokenBinding.GetTotalSupply()
						if err != nil {
							zap.L().Error(
								"failed to get token total supply",
								zap.Error(err),
								zap.String("address", m.GetAddress().Hex()),
								zap.Any("network", report.Network),
							)
							m.results.Corrupted = true
						} else {
							m.results.TotalSupply = totalSupply
						}
					}
				}
			}
		}
	}

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TokenDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TokenDetector) Results() any {
	return m.results
}
