package utils

import "path/filepath"

// GetLocalSourcesPath returns the absolute path to the local sources directory.
func GetLocalSourcesPath() string {
	absPath, _ := filepath.Abs(filepath.Clean("../sources/"))
	return absPath
}
