package audit

// Report represents the top-level structure of the Slither JSON output.
type Report struct {
	Success bool     `json:"success"` // Indicates the success status of the audit.
	Error   *string  `json:"error"`   // Contains any error messages, if present.
	Results *Results `json:"results"` // Contains the results of the audit.
}

// Results encapsulates the list of detected vulnerabilities or issues.
type Results struct {
	Detectors []Detector `json:"detectors"` // List of detected vulnerabilities or issues.
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

// Element represents a specific element (e.g., function, contract) associated with a detected issue.
type Element struct {
	Type               string             `json:"type"`                        // Type of the element (e.g., "function", "contract").
	Name               string             `json:"name"`                        // Name of the element.
	SourceMapping      SourceMapping      `json:"source_mapping"`              // Source mapping details for the element.
	TypeSpecificFields TypeSpecificFields `json:"type_specific_fields"`        // Specific fields related to the element type.
	Signature          string             `json:"signature,omitempty"`         // Signature of the element, if applicable.
	AdditionalFields   *AdditionalFields  `json:"additional_fields,omitempty"` // Additional fields associated with the element.
}

// SourceMapping provides details about the source code location of an element.
type SourceMapping struct {
	Start            int    `json:"start"`             // Start position in the source code.
	Length           int    `json:"length"`            // Length of the code segment.
	FilenameRelative string `json:"filename_relative"` // Relative path to the source file.
	FilenameAbsolute string `json:"filename_absolute"` // Absolute path to the source file.
	FilenameShort    string `json:"filename_short"`    // Short name of the source file.
	IsDependency     bool   `json:"is_dependency"`     // Indicates if the element is a dependency.
	Lines            []int  `json:"lines"`             // Line numbers associated with the code segment.
	StartingColumn   int    `json:"starting_column"`   // Starting column of the code segment.
	EndingColumn     int    `json:"ending_column"`     // Ending column of the code segment.
}

// TypeSpecificFields contains fields that are specific to the type of an element.
type TypeSpecificFields struct {
	Parent    *Element `json:"parent"`              // Parent element, if applicable.
	Directive []string `json:"directive,omitempty"` // Directive associated with the element, if applicable.
}

// AdditionalFields provides additional information about an element.
type AdditionalFields struct {
	UnderlyingType string `json:"underlying_type"`         // Underlying type of the element.
	VariableName   string `json:"variable_name,omitempty"` // Name of the variable, if applicable.
}
