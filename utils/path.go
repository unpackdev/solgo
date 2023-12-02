package utils

import (
	"os"
	"path/filepath"
)

// GetLocalSourcesPath returns the absolute path to the local sources directory.
func GetLocalSourcesPath() string {
	absPath, _ := filepath.Abs(filepath.Clean("../sources/"))
	return absPath
}

// PathExists returns true if the given path exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetCurrentPath returns the current working directory.
func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}
