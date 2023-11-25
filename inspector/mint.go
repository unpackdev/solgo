package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type MintResults struct {
	Enabled            bool              `json:"enabled"`
	Visibility         ast_pb.Visibility `json:"visibility"`
	ExternallyCallable bool              `json:"externally_callable"`
	Function           *ast.Function     `json:"function"`
}

func (m MintResults) IsEnabled() bool {
	return m.Enabled
}

func (m MintResults) IsVisible() bool {
	return m.Visibility == ast_pb.Visibility_PUBLIC || m.Visibility == ast_pb.Visibility_EXTERNAL
}

type MintDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *MintResults
}

func NewMintDetector(ctx context.Context, inspector *Inspector) Detector {
	return &MintDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"mint", "mintFor", "mintTo", "mintWithTokenURI", "mintBatch", "mintBatchFor", "mintBatchTo", "mintBatchWithTokenURI",
			"_mint", "_mintFor", "_mintTo", "_mintWithTokenURI", "_mintBatch", "_mintBatchFor", "_mintBatchTo", "_mintBatchWithTokenURI",
		},
		results: &MintResults{},
	}
}

func (m *MintDetector) Name() string {
	return "Mint Detector"
}

func (m *MintDetector) Type() DetectorType {
	return MintDetectorType
}

func (m *MintDetector) RegisterFunctionName(fnName string) bool {
	if !utils.StringInSlice(fnName, m.functionNames) {
		m.functionNames = append(m.functionNames, fnName)
		return true
	}

	return false
}

func (m *MintDetector) GetFunctionNames() []string {
	return m.functionNames
}

func (m *MintDetector) FunctionNameExists(fnName string) bool {
	return utils.StringInSlice(fnName, m.functionNames)
}

// Enter for now does nothing for mint detector. It may be needed in the future.
func (m *MintDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *MintDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
			switch nodeCtx := node.(type) {
			case *ast.Constructor:
				//utils.DumpNodeNoExit(nodeCtx.Parameters)
			case *ast.Function:

				if m.FunctionNameExists(nodeCtx.GetName()) {
					m.results.Enabled = true
					m.results.Visibility = nodeCtx.GetVisibility()
					//m.results.ExternallyCallable = function.IsExternallyCallable()
					//m.results.Function = nodeCtx
				}
			}
			return true
		},
	}
}

func (m *MintDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {

			return true
		},
	}
}

func (m *MintDetector) Results() any {
	return m.results
}
