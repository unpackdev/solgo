package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContract represents a contract used for testing purposes
type TestContract struct {
	Path    string // File path of the contract
	Content string // Content of the contract file
	Bytes   []byte // Byte representation of the contract
}

// ReadContractFileForTest reads a contract file for testing
func ReadContractFileForTest(t *testing.T, name string) TestContract {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	contractsDir := filepath.Join(dir, "..", "data", "tests")
	path := filepath.Join(contractsDir, name+".sol")

	_, err = os.Stat(contractsDir)
	assert.NoError(t, err)

	content, err := os.ReadFile(filepath.Clean(path))
	assert.NoError(t, err)

	return TestContract{Path: path, Content: string(content)}
}

// ReadContractFileForTestFromRootPath reads a contract file for testing from root of the solgo project
func ReadContractFileForTestFromRootPath(t *testing.T, name string) TestContract {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fmt.Println(dir)

	contractsDir := filepath.Join(dir, "data", "tests")
	path := filepath.Join(contractsDir, name+".sol")

	_, err = os.Stat(contractsDir)
	assert.NoError(t, err)

	content, err := os.ReadFile(filepath.Clean(path))
	assert.NoError(t, err)

	return TestContract{Path: path, Content: string(content)}
}

// ReadJsonBytesForTest reads a json file for testing purposes
func ReadJsonBytesForTest(t *testing.T, name string) TestContract {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	contractsDir := filepath.Join(dir, "..", "data", "tests")
	path := filepath.Join(contractsDir, name+".json")

	_, err = os.Stat(contractsDir)
	assert.NoError(t, err)

	content, err := os.ReadFile(filepath.Clean(path))
	assert.NoError(t, err)

	return TestContract{Path: path, Content: string(content), Bytes: content}
}

// ReadJsonBytesForTestFromRootPath reads a json file for testing purposes from root of the solgo project
func ReadJsonBytesForTestFromRootPath(t *testing.T, name string) TestContract {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	contractsDir := filepath.Join(dir, "data", "tests")
	path := filepath.Join(contractsDir, name+".json")

	_, err = os.Stat(contractsDir)
	assert.NoError(t, err)

	content, err := os.ReadFile(filepath.Clean(path))
	assert.NoError(t, err)

	return TestContract{Path: path, Content: string(content), Bytes: content}
}
