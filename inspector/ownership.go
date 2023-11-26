package inspector

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type OwnershipResults struct {
	Detected bool `json:"detected"`
}

type OwnershipDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *OwnershipResults
}

func NewOwnershipDetector(ctx context.Context, inspector *Inspector) Detector {
	return &OwnershipDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"transferOwnership", "renounceOwnership", "_transferOwnership", "_renounceOwnership",
			"owner", "isOwner", "setOwner", "claimOwnership", "initializeOwnership",
		},
		results: &OwnershipResults{},
	}
}

func (m *OwnershipDetector) Name() string {
	return "State Variable Detector"
}

func (m *OwnershipDetector) Type() DetectorType {
	return OwnershipDetectorType
}

func (m *OwnershipDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fn.GetName(), m.functionNames) {
					m.results.Detected = true
					fmt.Printf("Detected ownership function: %s\n", fn.GetName())
				}
			}
			return true, nil
		},
	}
}

func (m *OwnershipDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *OwnershipDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *OwnershipDetector) Results() any {
	return m.results
}
