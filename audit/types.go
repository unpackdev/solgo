package audit

import "encoding/json"

type Response struct {
	Success bool     `json:"success"`
	Error   *string  `json:"error"`
	Results *Results `json:"results"`
}

type Results struct {
	Detectors []Detector `json:"detectors"`
}

type Detector struct {
	Elements             []Element `json:"elements"`
	Description          string    `json:"description"`
	Markdown             string    `json:"markdown"`
	FirstMarkdownElement string    `json:"first_markdown_element"`
	ID                   string    `json:"id"`
	Check                string    `json:"check"`
	Impact               string    `json:"impact"`
	Confidence           string    `json:"confidence"`
}

type Element struct {
	Type               string             `json:"type"`
	Name               string             `json:"name"`
	SourceMapping      SourceMapping      `json:"source_mapping"`
	TypeSpecificFields TypeSpecificFields `json:"type_specific_fields"`
	Signature          string             `json:"signature,omitempty"`
	AdditionalFields   *AdditionalFields  `json:"additional_fields,omitempty"`
}

type SourceMapping struct {
	Start            int    `json:"start"`
	Length           int    `json:"length"`
	FilenameRelative string `json:"filename_relative"`
	FilenameAbsolute string `json:"filename_absolute"`
	FilenameShort    string `json:"filename_short"`
	IsDependency     bool   `json:"is_dependency"`
	Lines            []int  `json:"lines"`
	StartingColumn   int    `json:"starting_column"`
	EndingColumn     int    `json:"ending_column"`
}

type TypeSpecificFields struct {
	Parent    *Element `json:"parent"`
	Directive []string `json:"directive,omitempty"`
}

type AdditionalFields struct {
	UnderlyingType string `json:"underlying_type"`
	VariableName   string `json:"variable_name,omitempty"`
}

// NewResponse creates a new Response object from the given data in bytes.
// Bytes will be Slither JSON output.
// Returns an error if the given data is not a valid JSON.
func NewResponse(data []byte) (*Response, error) {
	var toReturn *Response
	if err := json.Unmarshal(data, &toReturn); err != nil {
		return nil, err
	}
	return toReturn, nil
}
