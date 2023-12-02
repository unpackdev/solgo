package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/storage"
	"go.uber.org/zap"
)

type StorageResults struct {
	Detected   bool                `json:"detected"`
	Descriptor *storage.Descriptor `json:"descriptor"`
}

type StorageDetector struct {
	ctx context.Context
	*Inspector
	results *StorageResults
}

func NewStorageDetector(ctx context.Context, inspector *Inspector) Detector {
	return &StorageDetector{
		ctx:       ctx,
		Inspector: inspector,
		results:   &StorageResults{},
	}
}

func (m *StorageDetector) Name() string {
	return "Storage Detector"
}

func (m *StorageDetector) Type() DetectorType {
	return StorageDetectorType
}

func (m *StorageDetector) Enter(ctx context.Context) (DetectorFn, error) {
	// As of now, we do not need to traverse through the AST.
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *StorageDetector) Detect(ctx context.Context) (DetectorFn, error) {
	reader, err := m.GetStorage().Describe(ctx, m.GetAddress(), m.GetDetector(), nil)
	if err != nil {
		zap.L().Error(
			"failed to describe contract storage",
			zap.Error(err),
			zap.String("address", m.GetAddress().Hex()),
		)
		return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, err
	} else {
		m.results.Detected = true
		m.results.Descriptor = reader.GetDescriptor()
	}

	// As of now, we do not need to traverse through the AST.
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *StorageDetector) Exit(ctx context.Context) (DetectorFn, error) {

	// As of now, we do not need to traverse through the AST.
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *StorageDetector) Results() any {
	return m.results
}
