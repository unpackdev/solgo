package inspector

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

type Function struct {
	ContractName string                       `json:"contract_name"`
	ContractType ast_pb.NodeType              `json:"contract_type"`
	ContractKind ast_pb.NodeType              `json:"contract_kind"`
	Name         string                       `json:"name"`
	Standard     *standards.FunctionDiscovery `json:"standard"`
	Statement    ast.Node[ast_pb.NodeType]    `json:"statement"`
}

type TransferResults struct {
	Detected  bool       `json:"detected"`
	Safe      bool       `json:"safe"`
	Functions []Function `json:"functions"`
}

type TransferDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *TransferResults
}

func NewTransferDetector(ctx context.Context, inspector *Inspector) Detector {
	return &TransferDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"transfer", "transferFrom", "_transfer", "_transferFrom",
		},
		results: &TransferResults{
			Safe:      true,
			Functions: make([]Function, 0),
		},
	}
}

func (m *TransferDetector) Name() string {
	return "Transfer Detector"
}

func (m *TransferDetector) Type() DetectorType {
	return TransferDetectorType
}

func (m *TransferDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool {

	standard, err := standards.GetContractByStandard(standards.ERC20)
	if err != nil {
		fmt.Println(err)
	}

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
			if fnCtx, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fnCtx.GetName(), m.functionNames) {
					var discoveredFn Function
					discoveredFn.Name = fnCtx.GetName()

					if contract := m.inspector.GetDetector().GetAST().GetTree().GetById(fnCtx.GetScope()); contract != nil {
						discoveredFn.ContractType = contract.GetType()
						switch contractCtx := contract.(type) {
						case *ast.Contract:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						case *ast.Interface:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						case *ast.Library:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						}
					}

					if standardFn := m.getStandardFunction(standard, fnCtx.GetName()); standardFn != nil {
						m.results.Detected = true
						newStandardFn := m.buildStandardFunction(fnCtx)
						if check, found := standards.FunctionConfidenceCheck(standard, &newStandardFn); found {
							discoveredFn.Standard = &check
						}
					} else { // handle _transfer and _transferFrom

					}

					m.results.Functions = append(m.results.Functions, discoveredFn)
				}

			}
			return true
		},
	}
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

func (m *TransferDetector) getStandardFunction(standard standards.EIP, fnName string) *standards.Function {
	standardFunctions := standard.GetFunctions()
	for _, fn := range standardFunctions {
		if fn.Name == fnName {
			return &fn
		}
	}
	return nil
}

func (m *TransferDetector) buildStandardFunction(fnCtx *ast.Function) standards.Function {
	var inputs []standards.Input
	var outputs []standards.Output

	if parametersList := fnCtx.GetParameters(); parametersList != nil {
		if parameters := parametersList.GetParameters(); parameters != nil {
			for _, param := range parameters {
				inputs = append(inputs, standards.Input{
					Type: param.GetTypeName().GetName(),
				})
			}
		}
	}

	if returnsList := fnCtx.GetReturnParameters(); returnsList != nil {
		if returns := returnsList.GetParameters(); returns != nil {
			for _, ret := range returns {
				outputs = append(outputs, standards.Output{
					Type: ret.GetTypeName().GetName(),
				})
			}
		}
	}

	return standards.Function{
		Name:    fnCtx.GetName(),
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func (m *TransferDetector) validateParameters(fn *standards.Function, fnCtx *ast.Function) bool {
	if parametersList := fnCtx.GetParameters(); parametersList != nil {
		if parameters := parametersList.GetParameters(); parameters != nil {
			if len(parameters) == len(fn.Inputs) {
				fmt.Println("Inputs count match...")
				//return true
			} else {
				fmt.Println("Inputs count mismatch...")
				//return false
			}

			for _, param := range parameters {
				for _, input := range fn.Inputs {
					fmt.Println("Parameter Type:", param.GetTypeName().GetName(), "Input Type:", input.Type)
					if param.GetName() == input.Type {
						fmt.Println("Parameter name match...")
					}
				}
			}

		}
	}

	return false
}
