package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ctxKeyRequestID int

const requestIDKey ctxKeyRequestID = 0

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func GetReqID(r *http.Request) string {
	reqID, _ := r.Context().Value(requestIDKey).(string)
	return reqID
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-Id")
		if reqID == "" {
			reqID = nextRequestID()
		}

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		w.Header().Set("X-Request-Id", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
