package inspector

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/standards"
	"go.uber.org/zap"
)

type Inspector struct {
	ctx       context.Context
	addresses []common.Address
	detector  *detector.Detector
	report    *Report
	visitor   *ast.NodeVisitor
}

func NewInspector(detector *detector.Detector, addresses ...common.Address) (*Inspector, error) {
	return &Inspector{
		addresses: addresses,
		detector:  detector,
		visitor:   &ast.NodeVisitor{},
		report: &Report{
			Addresses: addresses,
			Detectors: make(map[DetectorType]any),
		},
	}, nil
}

func (i *Inspector) GetAddresses() []common.Address {
	return i.addresses
}

func (i *Inspector) AddressExists(address common.Address) bool {
	for _, addr := range i.addresses {
		if addr == address {
			return true
		}
	}

	return false
}

func (i *Inspector) RegisterAddress(address common.Address) bool {
	if !i.AddressExists(address) {
		i.addresses = append(i.addresses, address)
		i.report.Addresses = append(i.report.Addresses, address)
		return true
	}

	return false
}

func (i *Inspector) GetDetector() *detector.Detector {
	return i.detector
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
	transferCheckFunc := func(node ast.Node[ast.NodeType]) bool {
		functionNode, ok := node.(*ast.Function)
		if !ok {
			return true // Not a function node, skip
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

		return true // Continue walking
	}

	i.detector.GetAST().GetTree().ExecuteTypeVisit(ast_pb.NodeType_FUNCTION_DEFINITION, transferCheckFunc)
	return i.report.UsesTransfers
}

func (i *Inspector) Inspect(only ...DetectorType) error {
	zap.L().Info("Inspecting contract")

	// Iterate through each registered detector and execute their logic
	for detectorType, detector := range registry {
		zap.L().Info("Running detector", zap.String("DetectorType", string(detectorType)))

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
