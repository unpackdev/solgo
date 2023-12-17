package inspector

import (
	"context"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type AntiWhaleResults struct {
	Detected bool                `json:"detected"`
	Provider utils.AntiWhaleType `json:"provider"`
}

type AntiWhaleDetector struct {
	ctx                context.Context
	inspector          *Inspector
	contractNames      []string
	stateVariableNames []string
	functionNames      []string
	results            *AntiWhaleResults
}

func NewAntiWhaleDetector(ctx context.Context, inspector *Inspector) Detector {
	return &AntiWhaleDetector{
		ctx:                ctx,
		inspector:          inspector,
		contractNames:      []string{"IPinkAntiBot"},
		stateVariableNames: []string{"antiBotEnabled"},
		functionNames:      []string{"setEnableAntiBot", "setTokenOwner", "setAntiBotEnabled", "onPreTransferCheck"},
		results:            &AntiWhaleResults{},
	}
}

func (m *AntiWhaleDetector) Name() string {
	return "External Calls Detector"
}

func (m *AntiWhaleDetector) Type() DetectorType {
	return AntiWhaleDetectorType
}

// SetInspector sets the inspector for the detector
func (m *AntiWhaleDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *AntiWhaleDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *AntiWhaleDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_MEMBER_ACCESS: func(node ast.Node[ast.NodeType]) (bool, error) {
			if ma, ok := node.(*ast.MemberAccessExpression); ok {
				if expr, ok := ma.GetExpression().(*ast.PrimaryExpression); ok {
					if expr.GetTypeDescription() != nil && strings.Contains(expr.GetTypeDescription().GetString(), "contract") {
						// This is basically for pinksale anti bot...
						// https://github.com/ctonydev/pink-antibot-guide-binance-solidity
						if utils.StringInSlice(ma.GetMemberName(), []string{"onPreTransferCheck"}) {
							m.results.Detected = true

							// We are going to have contract type in the type description
							if strings.Contains(expr.GetTypeDescription().GetString(), "IPinkAntiBot") {
								m.results.Provider = utils.AntiWhalePinksale
							}

							return false, nil
						}
					}
				}

			}
			return true, nil
		},
	}, nil
}

func (m *AntiWhaleDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AntiWhaleDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *AntiWhaleDetector) Results() any {
	return m.results
}
