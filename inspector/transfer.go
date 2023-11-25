package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type TransferResults struct{}

type TransferDetector struct {
	ctx       context.Context
	inspector *Inspector
	enabled   bool
	results   *TransferResults
}

func NewTransferDetector(ctx context.Context, inspector *Inspector) Detector {
	return &TransferDetector{
		ctx:       ctx,
		inspector: inspector,
		enabled:   false,
		results:   &TransferResults{},
	}
}

func (m *TransferDetector) Name() string {
	return "Transfer Detector"
}

func (m *TransferDetector) Type() DetectorType {
	return TransferDetectorType
}

func (m *TransferDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *TransferDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *TransferDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{}
}

func (m *TransferDetector) Results() any {
	return m.results
}
