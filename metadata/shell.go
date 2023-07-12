package metadata

import "io"

// Shell is an interface that wraps the Cat method from ipfs.Shell.
type Shell interface {
	Cat(path string) (io.ReadCloser, error)
}
