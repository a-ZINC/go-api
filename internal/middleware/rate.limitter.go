package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RateLimiter struct {
	mu sync.Mutex
	limit int
	visitors map[string]int
	duration time.Duration
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		mu:       sync.Mutex{},
		limit:    limit,
		visitors: make(map[string]int),
		duration: duration,
	}
	go rl.resetVisitorCounts()
	return rl
}
func (rl *RateLimiter) resetVisitorCounts() {
	for {
		time.Sleep(rl.duration)
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}
func (rl *RateLimiter) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		ip := rl.GetIP(r)
		rl.visitors[ip]++
		fmt.Printf("Visitor IP: %s, Count: %d\n", ip, rl.visitors[ip])
		if rl.visitors[ip] > rl.limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) GetIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}