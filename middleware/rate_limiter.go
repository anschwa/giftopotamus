package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/anschwa/giftopotamus/logger"
	"golang.org/x/time/rate"
)

type RateManager struct {
	mu       sync.Mutex
	requests map[string]*rate.Limiter
}

func NewRateManager() *RateManager {
	return &RateManager{
		mu:       sync.Mutex{},
		requests: make(map[string]*rate.Limiter),
	}
}

func (rm *RateManager) Limit(dt time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rm.mu.Lock()
			defer rm.mu.Unlock()

			reqID := GetReqID(r)

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				logger.Error(reqID, err)
				c := http.StatusInternalServerError
				http.Error(w, http.StatusText(c), c)
				return
			}

			limiter, ok := rm.requests[ip]
			if !ok {
				// Allow up to 3 requests per rate limit interval
				limiter = rate.NewLimiter(rate.Every(dt), 3)
				rm.requests[ip] = limiter
			}

			if !limiter.Allow() {
				c := http.StatusTooManyRequests
				http.Error(w, http.StatusText(c), c)
				return
			}

			// Handle request if rate limit has not been exceeded
			next.ServeHTTP(w, r)
		})
	}
}
