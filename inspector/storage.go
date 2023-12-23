package inspector

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/cfg"
	"github.com/unpackdev/solgo/storage"
	"go.uber.org/zap"
)

type StorageResults struct {
	Detected   bool                `json:"detected"`
	Descriptor *storage.Descriptor `json:"descriptor"`
	Cfg        *cfg.Builder        `json:"cfg"`
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

func (m *StorageDetector) GetResults() any {
	return m.results
}

// SetInspector sets the inspector for the detector
func (m *StorageDetector) SetInspector(inspector *Inspector) {
	m.Inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *StorageDetector) GetInspector() *Inspector {
	return m.Inspector
}

func (m *StorageDetector) Enter(ctx context.Context) (DetectorFn, error) {
	// As of now, we do not need to traverse through the AST.
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *StorageDetector) Detect(ctx context.Context) (DetectorFn, error) {
	cfgBuilder, err := cfg.NewBuilder(ctx, m.GetDetector().GetIR())
	if err != nil {
		return nil, fmt.Errorf("failed to create control flow graph builder: %w", err)
	}

	if err := cfgBuilder.Build(); err != nil {
		return nil, fmt.Errorf("failed to build control flow graph: %w", err)
	}
	m.results.Cfg = cfgBuilder

	reader, err := m.GetStorage().Describe(ctx, m.GetAddress(), m.GetDetector(), cfgBuilder, nil)
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
		//utils.DumpNodeNoExit(reader.GetDescriptor())
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
