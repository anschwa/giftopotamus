package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/anschwa/giftopotamus/logger"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			reqID := GetReqID(r)

			if r := recover(); r != nil {
				logger.Error(reqID, "PANIC RECOVERED", r, string(debug.Stack()))

				c := http.StatusInternalServerError
				http.Error(w, http.StatusText(c), c)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
