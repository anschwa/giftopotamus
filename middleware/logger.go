package middleware

import (
	"net/http"
	"time"

	"github.com/anschwa/giftopotamus/logger"
)

type StatusWriter struct {
	http.ResponseWriter
	Status int
}

func (r *StatusWriter) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &StatusWriter{ResponseWriter: w}

		t1 := time.Now()
		defer func() {
			reqID := GetReqID(r)
			logger.Info(reqID, sw.Status, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(t1))
		}()

		next.ServeHTTP(sw, r)
	})
}
