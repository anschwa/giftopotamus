package middleware

import (
	"context"
	"net/http"
	"time"
)

// Timeout wraps the next handler with an http.TimeoutHandler and sets
// the deadline on the request context.Context. This may seem
// redundant, but http.TimeoutHander will kill the connection to the
// client but the handler will keep executing on the server. This way
// each handler can decide to cancel execution of the request or not
// by selecting on the r.Context().Done() channel.
func Timeout(dt time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), dt)
			defer cancel()

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})

		return http.TimeoutHandler(h, dt, "Request timed out")
	}
}
