package solgo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	sources_pb "github.com/txpull/protos/dist/go/sources"
	"github.com/txpull/solgo/metadata"
	"github.com/txpull/solgo/utils"
)

var ErrPathFound = errors.New("path found")

// SourceUnit represents a unit of source code in Solidity. It includes the name, path, and content of the source code.
type SourceUnit struct {
	Name    string `yaml:"name" json:"name"`
	Path    string `yaml:"path" json:"path"`
	Content string `yaml:"content" json:"content"`
}

func (s *SourceUnit) String() string {
	return fmt.Sprintf("SourceUnit{Name: %s, Path: %s, Content: %s}", s.Name, s.Path, s.Content)
}

func (s *SourceUnit) GetName() string {
	return s.Name
}

func (s *SourceUnit) GetPath() string {
	return s.Path
}

func (s *SourceUnit) GetContent() string {
	return s.Content
}

// ToProto converts a SourceUnit to a protocol buffer SourceUnit.
func (s *SourceUnit) ToProto() *sources_pb.SourceUnit {
	return &sources_pb.SourceUnit{
		Name:    s.Name,
		Path:    s.Path,
		Content: s.Content,
	}
}

// Sources represents a collection of SourceUnit. It includes a slice of SourceUnit and the name of the entry source unit.
type Sources struct {
	prepared             bool          `yaml:"-" json:"-"`
	SourceUnits          []*SourceUnit `yaml:"source_units" json:"source_units"`
	EntrySourceUnitName  string        `yaml:"entry_source_unit" json:"base_source_unit"`
	MaskLocalSourcesPath bool          `yaml:"mask_local_sources_path" json:"mask_local_sources_path"`
	LocalSourcesPath     string        `yaml:"local_sources_path" json:"local_sources_path"`
}

// ArePrepared returns true if the Sources has been prepared.
func (s *Sources) ArePrepared() bool {
	return s.prepared
}

// GetUnits returns the SourceUnits in the Sources.
func (s *Sources) GetUnits() []*SourceUnit {
	return s.SourceUnits
}

// ToProto converts a Sources to a protocol buffer Sources.
func (s *Sources) ToProto() *sources_pb.Sources {
	var sourceUnits []*sources_pb.SourceUnit
	for _, sourceUnit := range s.SourceUnits {
		sourceUnits = append(sourceUnits, sourceUnit.ToProto())
	}

	return &sources_pb.Sources{
		EntrySourceUnitName:  s.EntrySourceUnitName,
		MaskLocalSourcesPath: s.MaskLocalSourcesPath,
		LocalSourcesPath:     s.LocalSourcesPath,
		SourceUnits:          sourceUnits,
	}
}

// NewSourcesFromMetadata creates a Sources from a metadata package ContractMetadata.
// This is a helper function that ensures easier integration when working with the metadata package.
func NewSourcesFromMetadata(md *metadata.ContractMetadata) *Sources {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	sourcesDir := filepath.Clean(filepath.Join(dir, "sources"))
	sources := &Sources{
		MaskLocalSourcesPath: true,
		LocalSourcesPath:     sourcesDir,
	}

	// First target is the target of the entry source unit...
	for _, name := range md.Settings.CompilationTarget {
		sources.EntrySourceUnitName = name
		break
	}

	// Getting name looks surreal and easy, probably won't work in all cases and is
	// too good to be true.
	for name, source := range md.Sources {
		sources.SourceUnits = append(sources.SourceUnits, &SourceUnit{
			Name:    strings.TrimSuffix(filepath.Base(name), ".sol"),
			Path:    name,
			Content: source.Content,
		})
	}

	return sources
}

// Validate checks the integrity of the Sources object.
// It ensures that:
// - There is at least one SourceUnit.
// - Each SourceUnit has a name and either a path or content.
// - If a SourceUnit has a path, the file at that path exists.
// - The entry source unit name is valid.
func (s *Sources) Validate() error {
	// Ensure there is at least one SourceUnit.
	if len(s.SourceUnits) == 0 {
		return errors.New("no source units found")
	}

	// Validate each SourceUnit.
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Name == "" {
			return errors.New("source unit must have a name")
		}
		if sourceUnit.Path == "" && sourceUnit.Content == "" {
			return fmt.Errorf("source unit %s must have either path or content", sourceUnit.Name)
		}
	}

	// Validate the entry source unit name.
	if s.EntrySourceUnitName != "" {
		entrySourceUnit := s.GetSourceUnitByName(s.EntrySourceUnitName)
		if entrySourceUnit == nil {
			return fmt.Errorf("entry source unit %s not found", s.EntrySourceUnitName)
		}
	}

	return nil
}

// Prepare validates and prepares the Sources. It checks if each SourceUnit has either a path or content and a name.
// If a SourceUnit has a path but no content, it reads the content from the file at the path.
func (s *Sources) Prepare() error {

	// We should verify that path can be discovered if local sources path is
	// provided.
	if s.LocalSourcesPath != "" {
		if _, err := os.Stat(s.LocalSourcesPath); err != nil {
			return fmt.Errorf("local sources path %s does not exist", s.LocalSourcesPath)
		}
	} else {
		s.LocalSourcesPath = "./sources/"
	}

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

		// Extract import statements as perhaps some of them can be found in
		// local sources path and need to be prepended to the sources.
		importUnits, err := s.handleImports(sourceUnit)
		if err != nil {
			return err
		}

		s.SourceUnits = append(s.SourceUnits, importUnits...)
	}

	if err := s.SortContracts(); err != nil {
		return fmt.Errorf("failure while doing topological contract sorting: %s", err.Error())
	}

	// Mark sources as prepared for future use.
	s.prepared = true

	if err := s.Validate(); err != nil {
		return err
	}

	return nil
}

// GetCombinedSource combines the content of all SourceUnits in the Sources into a single string, separated by two newlines.
func (s *Sources) GetCombinedSource() string {
	var builder strings.Builder
	for i, sourceUnit := range s.SourceUnits {
		if i > 0 {
			builder.WriteString("\n\n")
		}
		builder.WriteString(sourceUnit.Content)
	}
	return builder.String()
}

// GetSourceUnitByName returns the SourceUnit with the given name from the Sources. If no such SourceUnit exists, it returns nil.
func (s *Sources) GetSourceUnitByName(name string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Name == name {
			return sourceUnit
		}
	}
	return nil
}

// GetSourceUnitByPath returns the SourceUnit with the given path from the Sources. If no such SourceUnit exists, it returns nil.
func (s *Sources) GetSourceUnitByPath(path string) *SourceUnit {
	for _, sourceUnit := range s.SourceUnits {
		if sourceUnit.Path == path {
			return sourceUnit
		}
	}
	return nil
}

// GetLocalSource attempts to find a local source file that matches the given partial path.
// It searches relative to the provided path and returns a SourceUnit representing the found source.
// If no matching source is found, it returns nil.
//
// The function replaces any instance of "@openzeppelin" in the partial path with the actual path to the openzeppelin-contracts repository.
// It then walks the file tree starting from "./sources/", checking each file against the new path.
//
// If the new path contains "../", it removes this and looks for the file in the parent directory.
// If a match is found, it creates a new SourceUnit with the name and path of the file, and returns it.
//
// If no "../" is present in the new path, it simply creates a new SourceUnit with the name and path.
//
// After a SourceUnit is created, the function checks if the file at the path exists.
// If it does, it reads the file content and assigns it to the SourceUnit's Content field.
// If the file does not exist, it returns an error.
//
// If the walk function encounters an error other than ErrPathFound, it returns the error.
// If the source is still nil after the walk, it returns nil.
func (s *Sources) GetLocalSource(partialPath string, relativeTo string) (*SourceUnit, error) {
	// Replace @openzeppelin with the actual path to the openzeppelin-contracts repository
	partialPath = replaceOpenZeppelin(partialPath)
	relativeTo = replaceOpenZeppelin(relativeTo)
	var source *SourceUnit
	errWalk := filepath.Walk(s.LocalSourcesPath, func(partialWalkPath string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		relativeToDir := filepath.Dir(relativeTo)
		newPath := filepath.Join(relativeToDir, partialPath)

		// If file contains ../, remove it and look for the file in the parent directory
		if strings.Contains(newPath, "../") {
			newPath = strings.TrimSpace(strings.Replace(newPath, "../", "", -1))
			if strings.Contains(partialWalkPath, newPath) {
				sourceName := strings.TrimSuffix(filepath.Base(newPath), ".sol")
				if !s.SourceUnitExists(sourceName) {
					source = &SourceUnit{
						Name: sourceName,
						Path: partialWalkPath,
					}
				}
				return ErrPathFound
			}
		}

		sourceName := strings.TrimSuffix(filepath.Base(newPath), ".sol")
		if !s.SourceUnitExists(sourceName) {
			if strings.Contains(partialWalkPath, newPath) {
				source = &SourceUnit{
					Name: sourceName,
					Path: partialWalkPath,
				}
			}
		}

		return nil
	})

	if errWalk != nil && errWalk != ErrPathFound {
		return nil, errWalk
	}

	if source == nil {
		return nil, nil
	}

	if _, err := os.Stat(source.Path); os.IsNotExist(err) {
		return nil, err
	}

	content, err := os.ReadFile(source.Path)
	if err != nil {
		return nil, err
	}

	source.Content = string(content)
	return source, nil
}

// SourceUnitExists returns true if a SourceUnit with the given name exists in the Sources.
func (s *Sources) SourceUnitExists(name string) bool {
	return s.SourceUnitExistsIn(name, s.SourceUnits)
}

// SourceUnitExistsIn returns true if a SourceUnit with the given name exists in the given slice of SourceUnits.
func (s *Sources) SourceUnitExistsIn(name string, units []*SourceUnit) bool {
	for _, sourceUnit := range units {
		if sourceUnit.Name == name {
			return true
		}
	}
	return false
}

// WriteToDir writes each SourceUnit's content to a file in the specified directory.
func (s *Sources) WriteToDir(path string) error {
	// Ensure the specified directory exists or create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0700); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", path, err)
		}
	}

	// Write each SourceUnit's content to a file in the specified directory
	for _, sourceUnit := range s.SourceUnits {
		content := utils.SimplifyImportPaths(sourceUnit.Content)

		filePath := filepath.Join(path, sourceUnit.Name+".sol")
		if err := utils.WriteToFile(filePath, []byte(content)); err != nil {
			return fmt.Errorf("failed to write source unit %s to file: %v", sourceUnit.Name, err)
		}
	}

	return nil
}

// TruncateDir removes all files and subdirectories within the specified directory.
func (s *Sources) TruncateDir(path string) error {
	// Open the directory
	dir, err := os.Open(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("failed to open directory %s: %v", path, err)
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("failed to read directory entries for %s: %v", path, err)
	}

	// Iterate over each entry and remove it
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			// If the entry is a directory, recursively remove it
			err := os.RemoveAll(entryPath)
			if err != nil {
				return fmt.Errorf("failed to remove directory %s: %v", entryPath, err)
			}
		} else {
			// If the entry is a file, remove it
			err := os.Remove(entryPath)
			if err != nil {
				return fmt.Errorf("failed to remove file %s: %v", entryPath, err)
			}
		}
	}

	return nil
}

// GetSolidityVersion extracts the highest Solidity version from all source units.
func (s *Sources) GetSolidityVersion() (string, error) {
	// Use a regular expression to match the pragma solidity statement
	// This regex will match versions like ^0.x.x and extract only 0.x.x
	re := regexp.MustCompile(`pragma solidity\s*\^?(\d+\.\d+\.\d+);`)

	var highestVersion string

	for _, sourceUnit := range s.SourceUnits {
		match := re.FindStringSubmatch(sourceUnit.Content)

		if len(match) >= 2 {
			currentVersion := match[1]
			if compareVersions(currentVersion, highestVersion) > 0 {
				highestVersion = currentVersion
			}
		}
	}

	if highestVersion == "" {
		return "", fmt.Errorf("no solidity version found in any source unit")
	}

	return highestVersion, nil
}

// handleImports extracts import statements from the source unit and adds them to the sources.
func (s *Sources) handleImports(sourceUnit *SourceUnit) ([]*SourceUnit, error) {
	imports := extractImports(sourceUnit.Content)
	var sourceUnits []*SourceUnit

	for _, imp := range imports {
		baseName := filepath.Base(imp)

		if !s.SourceUnitExists(baseName) {
			source, err := s.GetLocalSource(imp, sourceUnit.Path)
			if err != nil {
				return nil, err
			}

			// Source may not be found and no errors and that's ok, however, we don't want to append
			// nil source to the sources.
			if source == nil {
				continue
			}

			if !s.SourceUnitExistsIn(source.Name, sourceUnits) {
				sourceUnits = append(sourceUnits, source)
			}

			subUnits, err := s.handleImports(source)
			if err != nil {
				return nil, err
			}

			for _, subUnit := range subUnits {
				if !s.SourceUnitExistsIn(subUnit.Name, sourceUnits) {
					sourceUnits = append(sourceUnits, subUnit)
				}
			}
		}
	}

	return sourceUnits, nil
}

// extractImports extracts import statements from the source unit.
func extractImports(content string) []string {
	re := regexp.MustCompile(`import "(.*?)";`)
	matches := re.FindAllStringSubmatch(content, -1)

	imports := make([]string, len(matches))
	for i, match := range matches {
		imports[i] = match[1]
	}

	return imports
}

// replaceOpenZeppelin replaces the @openzeppelin path with the actual path to the openzeppelin-contracts repository.
func replaceOpenZeppelin(path string) string {
	return strings.Replace(path, "@openzeppelin", filepath.Join("./sources/", "openzeppelin"), 1)
}

// Node represents a unit of source code in Solidity with its dependencies.
type Node struct {
	Name         string
	Dependencies []string
}

// SortContracts sorts the SourceUnits based on their dependencies.
func (s *Sources) SortContracts() error {
	var nodes []Node
	for _, sourceUnit := range s.SourceUnits {
		imports := extractImports(sourceUnit.Content)
		var dependencies []string
		for _, imp := range imports {
			baseName := filepath.Base(imp)
			dependencies = append(dependencies, strings.TrimSuffix(baseName, ".sol"))
		}
		nodes = append(nodes, Node{Name: sourceUnit.Name, Dependencies: dependencies})
	}

	sortedNames, err := topologicalSort(nodes)
	if err != nil {
		return err
	}

	var sortedSourceUnits []*SourceUnit
	for _, name := range sortedNames {
		if sourceUnit := s.GetSourceUnitByName(name); sourceUnit != nil {
			sortedSourceUnits = append(sortedSourceUnits, sourceUnit)
		}
	}

	s.SourceUnits = sortedSourceUnits
	return nil
}

// topologicalSort sorts the nodes based on their dependencies.
func topologicalSort(nodes []Node) ([]string, error) {
	var sorted []string
	visited := make(map[string]bool)
	var visit func(nodeName string) error
	visit = func(nodeName string) error {
		if visited[nodeName] {
			return nil
		}
		visited[nodeName] = true
		for _, node := range nodes {
			if node.Name == nodeName {
				for _, dep := range node.Dependencies {
					if err := visit(dep); err != nil {
						return err
					}
				}
				break
			}
		}
		sorted = append(sorted, nodeName)
		return nil
	}

	for _, node := range nodes {
		if !visited[node.Name] {
			if err := visit(node.Name); err != nil {
				return nil, err
			}
		}
	}

	return sorted, nil
}

// compareVersions compares two version strings and returns:
// -1 if v1 < v2
//
//	0 if v1 == v2
//	1 if v1 > v2
func compareVersions(v1, v2 string) int {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	// This is the first run to comparation and therefore, we should ensure it's setting it
	// to the highest version.
	if len(v2Parts) < 3 {
		return 1
	}

	for i := 0; i < 3; i++ {
		v1Int, _ := strconv.Atoi(v1Parts[i])
		v2Int, _ := strconv.Atoi(v2Parts[i])
		if v1Int < v2Int {
			return -1
		} else if v1Int > v2Int {
			return 1
		}
	}
	return 0
}
