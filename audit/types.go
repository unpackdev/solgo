package audit

import (
	audit_pb "github.com/unpackdev/protos/dist/go/audit"
)

// Report represents the top-level structure of the Slither JSON output.
type Report struct {
	Success bool     `json:"success"` // Indicates the success status of the audit.
	Error   string   `json:"error"`   // Contains any error messages, if present.
	Results *Results `json:"results"` // Contains the results of the audit.
}

// IsSuccess returns true if the vulnerability report was generated successfully.
func (r *Report) IsSuccess() bool {
	return r.Success
}

// GetError returns the error message associated with the vulnerability report.
func (r *Report) GetError() string {
	return r.Error
}

// GetResults returns the Results struct associated with the vulnerability report.
func (r *Report) GetResults() *Results {
	return r.Results
}

// ToProto converts the Report struct to its protobuf representation.
func (r *Report) ToProto() *audit_pb.Report {
	return &audit_pb.Report{
		Success: r.Success,
		Error:   r.Error,
		Results: r.Results.ToProto(),
	}
}

// Results encapsulates the list of detected vulnerabilities or issues.
type Results struct {
	Detectors []Detector `json:"detectors"` // List of detected vulnerabilities or issues.
}

// GetDetectors returns the list of detected vulnerabilities or issues.
func (r *Results) GetDetectors() []Detector {
	return r.Detectors
}

// ToProto converts the Results struct to its protobuf representation.
func (r *Results) ToProto() *audit_pb.Results {
	var detectors []*audit_pb.Detector
	for _, d := range r.Detectors {
		detectors = append(detectors, d.ToProto())
	}
	return &audit_pb.Results{
		Detectors: detectors,
	}
}

// Detector represents a single detected vulnerability or issue.
type Detector struct {
	Elements             []Element `json:"elements"`               // Elements associated with the detected issue.
	Description          string    `json:"description"`            // Description of the detected issue.
	Markdown             string    `json:"markdown"`               // Markdown formatted description of the detected issue.
	FirstMarkdownElement string    `json:"first_markdown_element"` // The first markdown element related to the issue.
	ID                   string    `json:"id"`                     // Unique identifier for the detected issue.
	Check                string    `json:"check"`                  // The type or category of the detected issue.
	Impact               string    `json:"impact"`                 // The impact level of the detected issue.
	Confidence           string    `json:"confidence"`             // The confidence level of the detected issue.
}

// ToProto converts the Detector struct to its protobuf representation.
func (d *Detector) ToProto() *audit_pb.Detector {
	var elements []*audit_pb.Element
	for _, e := range d.Elements {
		elements = append(elements, e.ToProto())
	}
	return &audit_pb.Detector{
		Elements:             elements,
		Description:          d.Description,
		Markdown:             d.Markdown,
		FirstMarkdownElement: d.FirstMarkdownElement,
		Id:                   d.ID,
		Check:                d.Check,
		Impact:               d.Impact,
		Confidence:           d.Confidence,
	}
}

// Element represents a specific element (e.g., function, contract) associated with a detected issue.
type Element struct {
	Type               string             `json:"type"`                        // Type of the element (e.g., "function", "contract").
	Name               string             `json:"name"`                        // Name of the element.
	SourceMapping      SourceMapping      `json:"source_mapping"`              // Source mapping details for the element.
	TypeSpecificFields TypeSpecificFields `json:"type_specific_fields"`        // Specific fields related to the element type.
	Signature          string             `json:"signature,omitempty"`         // Signature of the element, if applicable.
	AdditionalFields   *AdditionalFields  `json:"additional_fields,omitempty"` // Additional fields associated with the element.
}

// ToProto converts the Element struct to its protobuf representation.
func (e *Element) ToProto() *audit_pb.Element {
	toReturn := &audit_pb.Element{
		Type:               e.Type,
		Name:               e.Name,
		SourceMapping:      e.SourceMapping.ToProto(),
		TypeSpecificFields: e.TypeSpecificFields.ToProto(),
		Signature:          e.Signature,
	}

	if e.AdditionalFields != nil {
		toReturn.AdditionalFields = e.AdditionalFields.ToProto()
	}

	return toReturn
}

// SourceMapping provides details about the source code location of an element.
type SourceMapping struct {
	Start            int     `json:"start"`             // Start position in the source code.
	Length           int     `json:"length"`            // Length of the code segment.
	FilenameRelative string  `json:"filename_relative"` // Relative path to the source file.
	FilenameAbsolute string  `json:"filename_absolute"` // Absolute path to the source file.
	FilenameShort    string  `json:"filename_short"`    // Short name of the source file.
	IsDependency     bool    `json:"is_dependency"`     // Indicates if the element is a dependency.
	Lines            []int32 `json:"lines"`             // Line numbers associated with the code segment.
	StartingColumn   int     `json:"starting_column"`   // Starting column of the code segment.
	EndingColumn     int     `json:"ending_column"`     // Ending column of the code segment.
}

// ToProto converts the SourceMapping struct to its protobuf representation.
func (sm *SourceMapping) ToProto() *audit_pb.SourceMapping {
	return &audit_pb.SourceMapping{
		Start:            int32(sm.Start),
		Length:           int32(sm.Length),
		FilenameRelative: sm.FilenameRelative,
		FilenameAbsolute: sm.FilenameAbsolute,
		FilenameShort:    sm.FilenameShort,
		IsDependency:     sm.IsDependency,
		Lines:            sm.Lines,
		StartingColumn:   int32(sm.StartingColumn),
		EndingColumn:     int32(sm.EndingColumn),
	}
}

// TypeSpecificFields contains fields that are specific to the type of an element.
type TypeSpecificFields struct {
	Parent    *Element `json:"parent"`              // Parent element, if applicable.
	Directive []string `json:"directive,omitempty"` // Directive associated with the element, if applicable.
}

// ToProto converts the TypeSpecificFields struct to its protobuf representation.
func (tsf *TypeSpecificFields) ToProto() *audit_pb.TypeSpecificFields {
	toReturn := &audit_pb.TypeSpecificFields{
		Directive: tsf.Directive,
	}

	if tsf.Parent != nil {
		toReturn.Parent = tsf.Parent.ToProto()
	}

	return toReturn
}

// AdditionalFields provides additional information about an element.
type AdditionalFields struct {
	UnderlyingType string `json:"underlying_type"`         // Underlying type of the element.
	VariableName   string `json:"variable_name,omitempty"` // Name of the variable, if applicable.
}

// ToProto converts the AdditionalFields struct to its protobuf representation.
func (af *AdditionalFields) ToProto() *audit_pb.AdditionalFields {
	return &audit_pb.AdditionalFields{
		UnderlyingType: af.UnderlyingType,
		VariableName:   af.VariableName,
	}
}
