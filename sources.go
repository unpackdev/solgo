package solgo

import (
	"fmt"
	"os"
	"strings"
)

// SourceUnit represents a unit of source code in Solidity. It includes the name, path, and content of the source code.
type SourceUnit struct {
	Name    string `yaml:"name" json:"name"`
	Path    string `yaml:"path" json:"path"`
	Content string `yaml:"content" json:"content"`
}

// Sources represents a collection of SourceUnit. It includes a slice of SourceUnit and the name of the entry source unit.
type Sources struct {
	SourceUnits         []SourceUnit `yaml:"source_units" json:"source_units"`
	EntrySourceUnitName string       `yaml:"entry_source_unit" json:"base_source_unit"`
}

// Prepare validates and prepares the Sources. It checks if each SourceUnit has either a path or content and a name.
// If a SourceUnit has a path but no content, it reads the content from the file at the path.
func (s Sources) Prepare() error {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Path == "" && sourceUnit.Content == "" {
			return fmt.Errorf("source unit must have either path or content")
		}

		if sourceUnit.Name == "" {
			return fmt.Errorf("source unit must have a name")
		}

		if sourceUnit.Path != "" && sourceUnit.Content == "" {
			content, err := os.ReadFile(sourceUnit.Path)
			if err != nil {
				return err
			}
			sourceUnit.Content = string(content)
		}
	}

	return nil
}

// GetCombinedSource combines the content of all SourceUnits in the Sources into a single string, separated by two newlines.
func (s Sources) GetCombinedSource() string {
	var sources []string
	for _, sourceUnit := range s.SourceUnits {
		sources = append(sources, sourceUnit.Content)
	}

	return strings.Join(sources, "\n\n")
}

// GetSourceUnitByName returns the SourceUnit with the given name from the Sources. If no such SourceUnit exists, it returns nil.
func (s Sources) GetSourceUnitByName(name string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Name == name {
			return &sourceUnit
		}
	}
	return nil
}

// GetSourceUnitByPath returns the SourceUnit with the given path from the Sources. If no such SourceUnit exists, it returns nil.
func (s Sources) GetSourceUnitByPath(path string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Path == path {
			return &sourceUnit
		}
	}
	return nil
}
