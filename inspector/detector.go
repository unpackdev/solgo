package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

type Detector interface {
	Name() string
	Type() DetectorType
	Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error)
	Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error)
	Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error)

	// // We are not able to use generics yet to the way I want to use them... Once it's enabled lets use it!
	// Basically we would need to use something like DetectorInterface but we cannot use it on registry variable declaration
	// due to compiler complaining errors.
	Results() any
}

func ToDetector[T any](d Detector) T {
	return d.(T)
}

func ToResults[T any](r any) T {
	return r.(T)
}

type DetectorResult struct {
	DetectionType       DetectionType          `json:"detection_type"`
	SeverityType        SeverityType           `json:"severity_type"`
	ConfidenceLevelType ConfidenceLevelType    `json:"detection_level"`
	Statement           ast.Node[ast.NodeType] `json:"statement"`
	SubDetectors        []DetectorResult       `json:"sub_detectors"`
}
