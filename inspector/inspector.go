package inspector

import (
	"github.com/ethereum/go-ethereum/common"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/detector"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

var (
	mintableFunctions = []string{
		"mint", "mintWithTokenURI", "mintBatch", "mintBatchWithTokenURI",
		"_mint", "_mintWithTokenURI", "_mintBatch", "_mintBatchWithTokenURI",
	}
)

type Inspector struct {
	addresses []common.Address
	detector  *detector.Detector
	report    *Report
}

func NewInspector(detector *detector.Detector, addresses ...common.Address) (*Inspector, error) {
	return &Inspector{
		addresses: addresses,
		detector:  detector,
		report: &Report{
			Addresses:      addresses,
			StateVariables: make([]StateVariable, 0),
			Transfers:      make([]Transfer, 0),
			Mintable: Mintable{
				Enabled: false,
			},
			Burnable: Burnable{
				Enabled: false,
			},
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
	return i.detector.GetIR().GetRoot().HasHighConfidenceStandard(standard)
}

func (i *Inspector) UsesTransfers() bool {
	var usesTransfers bool

	i.detector.GetIR().GetRoot().Walk(ast.NodeVisitor{
		TypeVisit: map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
			ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
				switch nodeCtx := node.(type) {
				case *ast.Function:
					if nodeCtx.GetName() == "transfer" || nodeCtx.GetName() == "transferFrom" {
						i.report.UsesTransfers = true
					}

					i.detector.GetAST().GetTree().WalkNodes(nodeCtx.GetNodes(), ast.NodeVisitor{
						TypeVisit: map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
							ast_pb.NodeType_FUNCTION_CALL: func(node ast.Node[ast.NodeType]) bool {
								subCtx := node.(*ast.FunctionCall)
								if subCtx.GetExpression() != nil && subCtx.GetExpression().GetType() == ast_pb.NodeType_MEMBER_ACCESS {
									expr := subCtx.GetExpression().(*ast.MemberAccessExpression)
									if expr.GetMemberName() == "transfer" || expr.GetMemberName() == "transferFrom" {
										usesTransfers = true
										i.report.UsesTransfers = true

										/* i.report.Transfers = append(i.report.Transfers, Transfer{
											Function:   nodeCtx,
											Expression: expr,
										}) */
									}
								}

								return true
							},
						},
					})
				}

				return true // Continue walking
			},
		},
	})

	return usesTransfers
}

func (i *Inspector) Inspect() error {
	zap.L().Info("Inspecting contract")

	i.detector.GetIR().GetRoot().Walk(ast.NodeVisitor{
		TypeVisit: map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
			ast_pb.NodeType_VARIABLE_DECLARATION: func(node ast.Node[ast.NodeType]) bool {

				switch nodeCtx := node.(type) {
				case *ast.VariableDeclaration:
					_ = nodeCtx
					/* 					i.report.StateVariables = append(i.report.StateVariables, StateVariable{
						Variable: nodeCtx,
					}) */
				}

				return true
			},
			ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
				switch nodeCtx := node.(type) {
				case *ast.Constructor:
					//utils.DumpNodeNoExit(nodeCtx.Parameters)
				case *ast.Function:

					// Lets detect if contract use mintable functionality...
					// If it's already discovered do not do it again....
					// @TODO: Figure out places where this function is being called out....
					if !i.report.Mintable.Enabled {
						if utils.StringInSlice(nodeCtx.GetName(), mintableFunctions) {
							i.report.Mintable.Enabled = true
							i.report.Mintable.Visibility = nodeCtx.GetVisibility()

							//i.report.Mintable.Function = nodeCtx

							utils.DumpNodeWithExit(i.report)
						}
					}

					// Lets detect if contract use burnable functionality...

				}

				return true
			},
		},
	})

	return i.resolve()
}

func (i *Inspector) resolve() error {
	return nil
}
