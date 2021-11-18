package handlers

import (
	"net/http"

	"github.com/anschwa/giftopotamus/middleware"
)

func SendGiftExchange(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorPage(w, http.StatusNotImplemented)
	})
}
