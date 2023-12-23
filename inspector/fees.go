package inspector

import (
	"context"
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

type FeesResults struct {
	Detected     bool `json:"detected"`
	CanModifyTax bool `json:"can_modify_tax"`
}

type FeesDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *FeesResults
}

func NewFeesDetector(ctx context.Context, inspector *Inspector) Detector {
	return &FeesDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"updatefees", "reducefee", "reducefees", "sendethtofee",
			"setfee", "setfeepercent", "setfees", "updatebuyfees", "updatesellfees",
			"setbuyfee", "setsellfee", "setbuyfees", "setsellfees",
			"settaxfeepercentown", "settaxfeepercent", "removeallfee", "removeallfees", "updatefee",
			"_settaxes", "", "restoreallfee", "restoreallfees", "totalbuytaxbasispoints", "totalselltaxbasispoints",
			"changefees", "togglecanswapfees", "changetax", "adjustfee", "setbtwtax", "_settax",
			"updatebuytaxes", "updateselltaxes", "setboosterfees", "setboosterfee", "setearlyselltax",
			"setearlybuytax", "setselltaxes", "updatebuytaxes", "handle_fees", "setservicefee", "returntonormaltax",
			"updatenetworkfee", "setprivatesalefee", "setbuytaxes", "setselltaxes", "updatetax",
			"settaxtozero", "updatefeeparams", "editmarketfee", "setrevenuefee", "changefee", "tokenfee",
			"_fee", "_customfee", "settaxes", "setselltaxes", "setbuytaxes", "setupfee", "_settaxes",
		},
		results: &FeesResults{},
	}
}

func (m *FeesDetector) Name() string {
	return "Pausable Detector"
}

func (m *FeesDetector) Type() DetectorType {
	return FeeDetectorType
}

// SetInspector sets the inspector for the detector
func (m *FeesDetector) SetInspector(inspector *Inspector) {
	m.inspector = inspector
}

// GetInspector returns the inspector for the detector
func (m *FeesDetector) GetInspector() *Inspector {
	return m.inspector
}

func (m *FeesDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(strings.ToLower(fn.GetName()), m.functionNames) {
					m.results.Detected = true
					m.results.CanModifyTax = true
					return false, nil
				}
			}
			return true, nil
		},
		ast_pb.NodeType_FUNCTION_CALL: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fn, ok := node.(*ast.FunctionCall); ok {
				if expr, ok := fn.GetExpression().(*ast.PrimaryExpression); ok {
					if utils.StringInSlice(strings.ToLower(expr.GetName()), m.functionNames) {
						m.results.Detected = true
						m.results.CanModifyTax = true
						return false, nil
					}
				}
			}
			return true, nil
		},
	}, nil
}

func (m *FeesDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *FeesDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *FeesDetector) Results() any {
	return m.results
}
