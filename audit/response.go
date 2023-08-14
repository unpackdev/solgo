package audit

import "encoding/json"

// ImpactLevel represents the severity of a detected issue in the audit results.
type ImpactLevel string

// String returns the string representation of the ImpactLevel.
func (i ImpactLevel) String() string {
	return string(i)
}

// Predefined impact levels representing the severity of detected issues.
const (
	ImpactHigh   ImpactLevel = "High"          // Represents high severity issues.
	ImpactMedium ImpactLevel = "Medium"        // Represents medium severity issues.
	ImpactLow    ImpactLevel = "Low"           // Represents low severity issues.
	ImpactInfo   ImpactLevel = "Informational" // Represents informational findings.
)

// NewResponse parses the provided JSON data (typically from Slither) and returns
// a structured Response object. If the data is not valid JSON or does not match
// the expected structure, an error is returned.
func NewResponse(data []byte) (*Response, error) {
	var toReturn *Response
	if err := json.Unmarshal(data, &toReturn); err != nil {
		return nil, err
	}
	return toReturn, nil
}

// FilterDetectorsByImpact filters the audit results based on the specified impact level
// and returns a list of detectors that match the given level.
func (r *Response) FilterDetectorsByImpact(impact ImpactLevel) []Detector {
	var filtered []Detector
	for _, detector := range r.Results.Detectors {
		if ImpactLevel(detector.Impact) == impact {
			filtered = append(filtered, detector)
		}
	}
	return filtered
}

// HasError determines if the audit response contains any error messages.
func (r *Response) HasError() bool {
	return r.Error != nil && *r.Error != ""
}

// ElementsByType retrieves all elements of a specified type from the audit results.
func (r *Response) ElementsByType(elementType string) []Element {
	var elements []Element
	for _, detector := range r.Results.Detectors {
		for _, element := range detector.Elements {
			if element.Type == elementType {
				elements = append(elements, element)
			}
		}
	}
	return elements
}

// UniqueImpactLevels identifies and returns a list of unique impact levels present
// in the audit results.
func (r *Response) UniqueImpactLevels() []string {
	impactSet := make(map[string]struct{})
	for _, detector := range r.Results.Detectors {
		impactSet[detector.Impact] = struct{}{}
	}
	var impacts []string
	for impact := range impactSet {
		impacts = append(impacts, impact)
	}
	return impacts
}

// DetectorsByCheck filters the audit results based on a specified check type
// and returns a list of detectors that match the given check.
func (r *Response) DetectorsByCheck(checkType string) []Detector {
	var detectors []Detector
	for _, detector := range r.Results.Detectors {
		if detector.Check == checkType {
			detectors = append(detectors, detector)
		}
	}
	return detectors
}

// CountByImpactLevel counts the number of detectors for each impact level
// and returns a map of impact levels to their respective counts.
func (r *Response) CountByImpactLevel() map[ImpactLevel]int {
	countMap := make(map[ImpactLevel]int)
	for _, detector := range r.Results.Detectors {
		countMap[ImpactLevel(detector.Impact)]++
	}
	return countMap
}

// HighConfidenceDetectors filters the audit results to return only those detectors
// that have a high confidence level.
func (r *Response) HighConfidenceDetectors() []Detector {
	var detectors []Detector
	for _, detector := range r.Results.Detectors {
		if detector.Confidence == "High" {
			detectors = append(detectors, detector)
		}
	}
	return detectors
}

// HasIssues determines if the audit response contains any detected issues or vulnerabilities.
func (r *Response) HasIssues() bool {
	return r.Results != nil && len(r.Results.Detectors) > 0
}
