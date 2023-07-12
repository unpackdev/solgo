package metadata

import (
	"io"

	ipfs "github.com/ipfs/go-ipfs-api"
)

// MockShell is a mock implementation of ipfs.Shell
type MockShell struct {
	ipfs.Shell
	CatFunc func(path string) (io.ReadCloser, error)
}

func (m *MockShell) Cat(path string) (io.ReadCloser, error) {
	return m.CatFunc(path)
}
