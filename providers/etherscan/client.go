package etherscan

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/redis/go-redis/v9"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// EtherScanProvider represents the BscScan scanner provider.
type EtherScanProvider struct {
	ctx      context.Context
	opts     *Options
	cache    *redis.Client
	keyIndex int32 // Use int32 for atomic operations
}

// NewEtherScanProvider creates a new instance of EtherScanProvider with the provided API key and API URL.
func NewEtherScanProvider(ctx context.Context, cache *redis.Client, opts *Options) *EtherScanProvider {
	return &EtherScanProvider{
		ctx:      ctx,
		opts:     opts,
		cache:    cache,
		keyIndex: 0,
	}
}

func (e *EtherScanProvider) ProviderName() string {
	return e.opts.Provider.String()
}

func (e *EtherScanProvider) CacheKey(method string, path string) string {
	return fmt.Sprintf(
		"etherscan::%s::%s",
		method,
		path,
	)
}

// GetNextKey returns the next API key in a round-robin fashion.
func (e *EtherScanProvider) GetNextKey() string {
	// Atomically increment the keyIndex and get the next index
	nextIndex := atomic.AddInt32(&e.keyIndex, 1)

	// Ensure the index wraps around the length of the keys slice
	keyCount := int32(len(e.opts.Keys))
	selectedIndex := nextIndex % keyCount

	return e.opts.Keys[selectedIndex]
}
