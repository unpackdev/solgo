package solgo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContract represents a contract used for testing purposes
type TestContract struct {
	Path    string // File path of the contract
	Content string // Content of the contract file
}

// ReadContractFileForTest reads a contract file for testing
func ReadContractFileForTest(t *testing.T, name string) TestContract {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	contractsDir := filepath.Join(dir, "data", "tests")
	path := filepath.Join(contractsDir, name+".sol")

	_, err = os.Stat(contractsDir)
	assert.NoError(t, err)

	content, err := os.ReadFile(path)
	assert.NoError(t, err)

	return TestContract{
		Path:    path,
		Content: string(content),
	}
}
