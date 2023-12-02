package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/storage"
)

type VariableDeclaration struct {
	Name          string                  `json:"name"`
	StateVariable bool                    `json:"state_variable"`
	Constant      bool                    `json:"constant"`
	Statement     ast.Node[ast.NodeType]  `json:"statement"`
	StorageSlot   *storage.SlotDescriptor `json:"storage_slot"`
}

type StateVariableResults struct {
	Detected     bool                   `json:"detected"`
	Declarations []*VariableDeclaration `json:"declarations"`
}

type StateVariableDetector struct {
	ctx context.Context
	*Inspector
	enabled bool
	results *StateVariableResults
}

func NewStateVariableDetector(ctx context.Context, inspector *Inspector) Detector {
	return &StateVariableDetector{
		ctx:       ctx,
		Inspector: inspector,
		enabled:   false,
		results: &StateVariableResults{
			Declarations: make([]*VariableDeclaration, 0),
		},
	}
}

func (m *StateVariableDetector) Name() string {
	return "State Variable Detector"
}

func (m *StateVariableDetector) Type() DetectorType {
	return StateVariableDetectorType
}

func (m *StateVariableDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_VARIABLE_DECLARATION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if varCtx, ok := node.(*ast.StateVariableDeclaration); ok {

				variable := &VariableDeclaration{
					Name:          varCtx.GetName(),
					StateVariable: varCtx.IsStateVariable(),
					Constant:      varCtx.IsConstant(),
					Statement:     varCtx,
				}

				// Neat trick to see if storage inspector is enabled and if within we discovered slot for the variable...
				if detector, ok := m.GetReport().Detectors[StorageDetectorType]; ok {
					if storageResults, ok := detector.(*StorageResults); ok {
						for _, slot := range storageResults.Descriptor.GetSlots() {
							if slot.Name == variable.Name {
								variable.StorageSlot = slot
							}
						}
					}
				}

				m.results.Declarations = append(m.results.Declarations, variable)
			}

			return true, nil
		},
	}
}

func (m *StateVariableDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *StateVariableDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

func (m *StateVariableDetector) Results() any {
	return m.results
}
