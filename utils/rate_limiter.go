package utils

import (
	"sync"
	"time"
)

// RateLimiter implements a token bucket rate limiting algorithm. It is used to control
// how frequently events are allowed to happen. The RateLimiter allows a certain number of
// events to occur within a fixed time frame and refills the tokens at a constant interval.
type RateLimiter struct {
	mutex      sync.Mutex    // Protects access to all RateLimiter fields.
	tokens     int           // Current number of available tokens.
	capacity   int           // Maximum number of tokens that can accumulate in the bucket.
	refillTime time.Duration // Time interval at which tokens are refilled to capacity.
	lastRefill time.Time     // Timestamp of the last refill operation.
}

// NewRateLimiter initializes and returns a new RateLimiter with the specified capacity and refill time.
// The capacity parameter specifies the maximum number of tokens, and refillTime specifies how often
// the tokens are replenished. A debug message showing the capacity is printed during initialization.
func NewRateLimiter(capacity int, refillTime time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     capacity,
		capacity:   capacity,
		refillTime: refillTime,
		lastRefill: time.Now(),
	}
}

// Allow checks if at least one token is available. If a token is available, it consumes
// a token by decrementing the token count and returns true, indicating that the event
// is allowed to proceed. If no tokens are available, it returns false, indicating that
// the rate limit has been exceeded. Tokens are refilled based on the elapsed time
// since the last refill, up to the maximum capacity.
func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	rl.refill()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// WaitForToken blocks the caller until a token becomes available. If no tokens are
// available upon invoking this method, it calculates the time until the next token refill
// and sleeps for that duration. This method ensures that events respect the rate limit by
// waiting for permission to proceed rather than immediately returning false.
func (rl *RateLimiter) WaitForToken() {
	for !rl.Allow() {
		rl.mutex.Lock()
		timeToNextRefill := rl.refillTime - time.Since(rl.lastRefill)
		rl.mutex.Unlock()
		if timeToNextRefill > 0 {
			time.Sleep(timeToNextRefill)
		}
	}
}

// refill is a helper method that replenishes the tokens based on the time elapsed
// since the last refill operation. If the elapsed time since the last refill is greater
// than or equal to the refill interval, the token count is reset to its maximum capacity.
// This method is called internally before checking if an event is allowed to proceed.
func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)

	if elapsed >= rl.refillTime {
		rl.tokens = rl.capacity
		rl.lastRefill = now
	}
}
