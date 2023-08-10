package utils

import (
	"os"
)

// WriteToFile writes the provided data byte array to a file at the provided path.
func WriteToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}
