package utils

import (
	"regexp"
	"strings"
)

// SimplifyImportPaths simplifies the paths in import statements as file will already be present in the
// directory for future consumption and is rather corrupted for import paths to stay the same.
func SimplifyImportPaths(content string) string {
	re := regexp.MustCompile(`import ".*?([^/]+\.sol)";`)
	return re.ReplaceAllString(content, `import "./$1";`)
}

// StripImportPaths removes the import paths entirely from the content.
func StripImportPaths(content string) string {
	re := regexp.MustCompile(`import ".*?";`)
	return re.ReplaceAllString(content, "")
}

// StripExtraSPDXLines removes the extra SPDX lines from the content.
// This is used when passing combined source to the solc compiler as it will complain about the extra SPDX lines.
func StripExtraSPDXLines(content string) string {
	lines := strings.Split(content, "\n")
	foundSPDX := false
	result := []string{}

	for _, line := range lines {
		if strings.HasPrefix(line, "// SPDX") {
			if !foundSPDX {
				result = append(result, line)
				foundSPDX = true
			}
		} else {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
