package clients

import "errors"

// ErrClientURLNotSet is returned when the client URL in the configuration is not set.
var ErrClientURLNotSet = errors.New("configuration client URL not set")

// ErrConcurrentClientsNotSet is returned when the number of concurrent clients in the configuration is not set.
var ErrConcurrentClientsNotSet = errors.New("configuration amount of concurrent clients is not set")

// ErrOptionsNotSet is returned when the configuration options are not set.
var ErrOptionsNotSet = errors.New("configuration options not set")

// ErrNodesNotSet is returned when the configuration nodes are not set.
var ErrNodesNotSet = errors.New("configuration nodes not set")
