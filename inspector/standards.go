package inspector

import (
	"context"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
	"github.com/unpackdev/solgo/standards"
)

// StandardsResults encapsulates the results of a standards detection process.
// It includes information about detected standards, their types, and confidence levels.
type StandardsResults struct {
	Detected                        bool                 `json:"detected"`                           // Flag indicating if any standard was detected.
	StandardTypes                   []standards.Standard `json:"standard_types"`                     // List of detected standard types.
	Standards                       []*ir.Standard       `json:"standards"`                          // List of detected standards.
	HighConfidenceMatchStandards    []*ir.Standard       `json:"high_confidence_match_standards"`    // Standards detected with high confidence.
	PerfectConfidenceMatchStandards []*ir.Standard       `json:"perfect_confidence_match_standards"` // Standards detected with perfect confidence.
}

// StandardsDetector is responsible for detecting contract standards in smart contracts.
// It extends the Inspector to analyze contracts and detect various standards.
type StandardsDetector struct {
	ctx        context.Context   // Context for managing async operations.
	*Inspector                   // Embedded Inspector for contract analysis.
	results    *StandardsResults // Results of the standards detection process.
}

// NewStandardsDetector creates a new instance of StandardsDetector.
// Initializes the detector with a given Inspector instance and an empty StandardsResults.
func NewStandardsDetector(ctx context.Context, inspector *Inspector) Detector {
	return &StandardsDetector{
		ctx:       ctx,
		Inspector: inspector,
		results: &StandardsResults{
			StandardTypes:                   make([]standards.Standard, 0),
			Standards:                       make([]*ir.Standard, 0),
			HighConfidenceMatchStandards:    make([]*ir.Standard, 0),
			PerfectConfidenceMatchStandards: make([]*ir.Standard, 0),
		},
	}
}

// Name returns the name of the StandardsDetector.
func (m *StandardsDetector) Name() string {
	return "Contract Standards Detector"
}

// Type returns the detector type of the StandardsDetector.
func (m *StandardsDetector) Type() DetectorType {
	return StandardsDetectorType
}

// SetInspector sets the Inspector instance of the StandardsDetector.
func (m *StandardsDetector) SetInspector(inspector *Inspector) {
	m.Inspector = inspector
}

// GetInspector returns the Inspector instance of the StandardsDetector.
func (m *StandardsDetector) GetInspector() *Inspector {
	return m.Inspector
}

// Enter initializes the detection process, setting up any necessary state or configuration.
func (m *StandardsDetector) Enter(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

// Detect performs the actual detection of standards in the smart contract.
// It analyzes the contract's intermediate representation to identify various standards.
// The method updates the results with detected standards and their confidence levels.
func (m *StandardsDetector) Detect(ctx context.Context) (DetectorFn, error) {
	if m.GetDetector() != nil && m.GetDetector().GetIR() != nil && m.GetDetector().GetIR().GetRoot() != nil {
		m.results.Detected = true
		irRoot := m.GetDetector().GetIR().GetRoot()
		m.results.Standards = irRoot.GetStandards()
		for _, standard := range m.results.Standards {
			if standard.GetConfidence().Confidence == standards.PerfectConfidence {
				m.results.StandardTypes = append(m.results.StandardTypes, standard.GetStandard().Type)
				m.results.PerfectConfidenceMatchStandards = append(m.results.PerfectConfidenceMatchStandards, standard)
			} else if standard.GetConfidence().Confidence == standards.HighConfidence {
				m.results.StandardTypes = append(m.results.StandardTypes, standard.GetStandard().Type)
				m.results.HighConfidenceMatchStandards = append(m.results.HighConfidenceMatchStandards, standard)
			}
		}
	}

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

// Exit finalizes the detection process, cleaning up any state or resources used.
func (m *StandardsDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

// Results returns the detection results stored in the StandardsDetector.
func (m *StandardsDetector) Results() any {
	return m.results
}
