// internal/middleware/ratelimit_test.go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Mock handler to simulate the behavior of the handler chain
func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Test rate limiting with the RateLimiter middleware
func TestRateLimiter(t *testing.T) {
	// Create a new rate limiter that allows only 2 requests per second
	rl := NewRateLimiter(2)

	// Create a mock HTTP request to pass through the rate limiter
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Test first request - this should pass
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the first request succeeded
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Test second request - this should pass as well
	rr = httptest.NewRecorder()
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the second request succeeded
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Test third request - this should be blocked (rate limit exceeded)
	rr = httptest.NewRecorder()
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the third request was blocked
	if rr.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status %d but got %d", http.StatusTooManyRequests, rr.Code)
	}

	// Wait for the rate limiter window to reset (after 1 second)
	time.Sleep(time.Second)

	// Test the fourth request - this should pass now that the rate limit is reset
	rr = httptest.NewRecorder()
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the fourth request succeeded
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Test the fifth request for a different IP address - this should pass (rate limit applies per IP)
	rr = httptest.NewRecorder()
	req.RemoteAddr = "192.168.1.2" // Change IP address for new client
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the fifth request for a different IP succeeded
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d for a different IP", http.StatusOK, rr.Code)
	}

	// Test that after 1 second the rate limit window resets for the original IP again
	time.Sleep(time.Second)

	// Test a new request from the original IP after rate limit window reset
	rr = httptest.NewRecorder()
	req.RemoteAddr = "192.168.1.1" // Reset back to original IP
	rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)

	// Check if the new request succeeded
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d after reset", http.StatusOK, rr.Code)
	}

	// Test that the rate limit works for high concurrency requests (simultaneous requests)
	// Send multiple requests at the same time to test concurrency handling

	done := make(chan bool, 5) // Channel to synchronize goroutines
	for i := 0; i < 5; i++ {
		go func() {
			rr := httptest.NewRecorder()
			rl.Limit(http.HandlerFunc(mockHandler)).ServeHTTP(rr, req)
			if rr.Code == http.StatusTooManyRequests {
				done <- true
			} else {
				done <- false
			}
		}()
	}

	// Check if one of the requests returned StatusTooManyRequests due to rate limiting
	count := 0
	for i := 0; i < 5; i++ {
		if <-done {
			count++
		}
	}
	if count == 0 {
		t.Errorf("Expected at least one request to be rate-limited but none were")
	}
}
