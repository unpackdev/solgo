package metadata

import "errors"

var (
	ErrInvalidIpfsClient      = errors.New("invalid ipfs client provided")
	ErrIpfsClientNotAvailable = errors.New("ipfs client seems not to be available. please check your ipfs daemon")
)
