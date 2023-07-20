package solgo

import (
	"fmt"
	"os"
	"strings"
)

type SourceUnit struct {
	Name    string `yaml:"name" json:"name"`
	Path    string `yaml:"path" json:"path"`
	Content string `yaml:"content" json:"content"`
}

type Sources struct {
	SourceUnits         []SourceUnit `yaml:"source_units" json:"source_units"`
	EntrySourceUnitName string       `yaml:"entry_source_unit" json:"base_source_unit"`
}

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

func (s Sources) GetCombinedSource() string {
	var sources []string
	for _, sourceUnit := range s.SourceUnits {
		sources = append(sources, sourceUnit.Content)
	}

	return strings.Join(sources, "\n\n")
}

func (s Sources) GetSourceUnitByName(name string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Name == name {
			return &sourceUnit
		}
	}
	return nil
}

func (s Sources) GetSourceUnitByPath(path string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Path == path {
			return &sourceUnit
		}
	}
	return nil
}
