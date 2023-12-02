package inspector

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/bindings"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/simulator"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/storage"
)

type Inspector struct {
	ctx         context.Context
	detector    *detector.Detector
	storage     *storage.Storage
	bindManager *bindings.Manager
	sim         *simulator.Simulator
	report      *Report
	visitor     *ast.NodeVisitor
}

func NewInspector(ctx context.Context, detector *detector.Detector, simulator *simulator.Simulator, storage *storage.Storage, bindManager *bindings.Manager, addr common.Address) (*Inspector, error) {
	return &Inspector{
		ctx:         ctx,
		detector:    detector,
		storage:     storage,
		bindManager: bindManager,
		sim:         simulator,
		visitor:     &ast.NodeVisitor{},
		report: &Report{
			Address:   addr,
			Detectors: make(map[DetectorType]any),
		},
	}, nil
}

func (i *Inspector) GetAddress() common.Address {
	return i.report.Address
}

func (i *Inspector) GetDetector() *detector.Detector {
	return i.detector
}

func (i *Inspector) GetStorage() *storage.Storage {
	return i.storage
}

func (i *Inspector) GetBindingManager() *bindings.Manager {
	return i.bindManager
}

func (i *Inspector) GetReport() *Report {
	return i.report
}

func (i *Inspector) IsReady() bool {
	return i.detector != nil && i.detector.GetIR() != nil && i.detector.GetIR().GetRoot() != nil && i.detector.GetIR().GetRoot().HasContracts()
}

func (i *Inspector) HasStandard(standard standards.Standard) bool {
	return i.detector.GetIR().GetRoot().HasHighConfidenceStandard(standard) || i.detector.GetIR().GetRoot().HasPerfectConfidenceStandard(standard)
}

func (i *Inspector) GetTree() *ast.Tree {
	return i.detector.GetAST().GetTree()
}

func (i *Inspector) UsesTransfers() bool {
	transferCheckFunc := func(node ast.Node[ast.NodeType]) (bool, error) {
		functionNode, ok := node.(*ast.Function)
		if !ok {
			return true, fmt.Errorf("node is not a function")
		}

		if functionNode.GetName() == "transfer" || functionNode.GetName() == "transferFrom" {
			i.report.UsesTransfers = true
		}

		for _, childNode := range functionNode.GetNodes() {
			functionCallNode, ok := childNode.(*ast.FunctionCall)
			if !ok {
				continue // Not a function call, skip
			}

			if expr := functionCallNode.GetExpression(); expr != nil && expr.GetType() == ast_pb.NodeType_MEMBER_ACCESS {
				memberAccessExpr, ok := expr.(*ast.MemberAccessExpression)
				if !ok {
					continue // Not a member access expression, skip
				}

				if memberAccessExpr.GetMemberName() == "transfer" || memberAccessExpr.GetMemberName() == "transferFrom" {
					i.report.UsesTransfers = true
				}
			}
		}

		return true, nil
	}

	i.detector.GetAST().GetTree().ExecuteTypeVisit(ast_pb.NodeType_FUNCTION_DEFINITION, transferCheckFunc)
	return i.report.UsesTransfers
}

func (i *Inspector) Inspect(only ...DetectorType) error {

	// Iterate through each registered detector and execute their logic
	for detectorType, detector := range registry {
		// If only is not empty, check if detector type is in only slice, if not continue to next detector
		if len(only) > 0 {
			if !IsDetectorType(detectorType, only...) {
				continue
			}
		}

		// Enter phase
		enterFuncs := detector.Enter(i.ctx)
		for nodeType, visitFunc := range enterFuncs {
			i.detector.GetAST().GetTree().ExecuteTypeVisit(nodeType, visitFunc)
		}

		// Detect phase
		detectFuncs := detector.Detect(i.ctx)
		for nodeType, visitFunc := range detectFuncs {
			i.detector.GetAST().GetTree().ExecuteTypeVisit(nodeType, visitFunc)
		}

		// Exit phase
		exitFuncs := detector.Exit(i.ctx)
		for nodeType, visitFunc := range exitFuncs {
			i.detector.GetAST().GetTree().ExecuteTypeVisit(nodeType, visitFunc)
		}

		results := detector.Results()
		i.report.Detectors[detectorType] = results
	}

	return i.resolve()
}

func (i *Inspector) resolve() error {
	return nil
}
