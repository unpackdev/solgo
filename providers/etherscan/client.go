package etherscan

import (
	"context"
	"fmt"
	"github.com/unpackdev/solgo/utils"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrorResponse represents the standard error response format returned by Etherscan's API.
type ErrorResponse struct {
	Status  string `json:"status"`  // The status code of the response.
	Message string `json:"message"` // A message describing the error.
	Result  string `json:"result"`  // The result field, typically empty in error responses.
}

// Provider encapsulates the logic for interacting with the Etherscan API,
// including handling API keys and caching responses.
type Provider struct {
	ctx         context.Context    // The context for controlling cancellations and timeouts.
	opts        *Options           // The configuration options for the provider.
	cache       *redis.Client      // A Redis client for caching API responses.
	keyIndex    int32              // An atomic counter for round-robin API key selection.
	rateLimiter *utils.RateLimiter // Helper to reduce lookups onto the etherscan a.k.a rate limit based on a subscription type.
}

// NewProvider initializes a new EtherScanProvider with specified options and cache.
// It returns an error if the provided options are invalid or incomplete.
func NewProvider(ctx context.Context, cache *redis.Client, opts *Options) (*Provider, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Provider{
		ctx:         ctx,
		opts:        opts,
		cache:       cache,
		keyIndex:    0,
		rateLimiter: utils.NewRateLimiter(opts.RateLimit*len(opts.Keys), 1*time.Second),
	}, nil
}

// ProviderName returns the name of the provider as specified in the options.
func (e *Provider) ProviderName() string {
	return e.opts.Provider.String()
}

// GetRateLimiter returns the instantiated rate limiter
func (e *Provider) GetRateLimiter() *utils.RateLimiter {
	return e.rateLimiter
}

// CacheKey generates a unique cache key for storing and retrieving API responses.
// The key is composed using the API method and path.
func (e *Provider) CacheKey(method string, path string) string {
	return fmt.Sprintf(
		"etherscan::%s::%s",
		method,
		path,
	)
}

// GetNextKey selects the next API key to use for a request in a round-robin fashion.
// This method ensures even distribution of request load across all configured API keys.
func (e *Provider) GetNextKey() string {
	// Atomically increment the keyIndex and get the next index
	nextIndex := atomic.AddInt32(&e.keyIndex, 1)

	// Ensure the index wraps around the length of the keys slice
	keyCount := int32(len(e.opts.Keys))
	selectedIndex := nextIndex % keyCount

	return e.opts.Keys[selectedIndex]
}
