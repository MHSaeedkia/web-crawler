// internal/middleware/ratelimit.go
package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"go.uber.org/ratelimit"
)

// RateLimiter is a struct that holds the ratelimiter instance and a request counter.
type RateLimiter struct {
	limiter       ratelimit.Limiter
	requestCounts map[string]int
	mu            sync.Mutex
	resetInterval time.Duration
}

// NewRateLimiter creates a new instance of RateLimiter with a given rate.
func NewRateLimiter(rate int) *RateLimiter {
	// Create a new rate limiter that allows 'rate' requests per second.
	limiter := ratelimit.New(rate)
	return &RateLimiter{
		limiter:       limiter,
		requestCounts: make(map[string]int),
		resetInterval: time.Second,
	}
}

// Limit is a middleware that enforces rate limiting on incoming requests.
func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		// Get the current timestamp in seconds
		currentTime := time.Now().Unix()

		// Use the IP address as the key for rate limiting
		clientIP := r.RemoteAddr

		// Check if the request count for the current second is within limits
		if rl.requestCounts[clientIP] >= 2 {
			// If the request count exceeds limit, send 429 Too Many Requests response
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			log.Println("Rate limit exceeded")
			return
		}

		// Increment the request count for this second
		rl.requestCounts[clientIP]++

		// Allow the request to pass through the rate limiter (take a token)
		rl.limiter.Take()

		// Log the allowed request
		log.Println("Request allowed")

		// Proceed to the next handler
		next.ServeHTTP(w, r)

		// Reset the counter for the current second after the interval
		go rl.resetRequestCount(clientIP, currentTime)
	})
}

// resetRequestCount resets the request count for the given client IP after the interval
func (rl *RateLimiter) resetRequestCount(clientIP string, currentTime int64) {
	// Wait for the reset interval before resetting the counter
	time.Sleep(rl.resetInterval)
	// Reset the request count for the given client
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Only reset if the second has passed
	if time.Now().Unix() > currentTime {
		delete(rl.requestCounts, clientIP)
	}
}
