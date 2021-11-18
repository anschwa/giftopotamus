package handlers

import (
	"net/http"

	"github.com/anschwa/giftopotamus/middleware"
)

func Logout(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/logout" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case "GET":
			http.Redirect(w, r, "/", http.StatusFound)
			return

		case "POST":
			sm.End(w, r)
			pd := &PageData{
				Title:      "Giftopotamus.com",
				SuccessMsg: "Successfully Logged out.",
			}

			tryRenderPage(w, r, PageLogout, pd)
			return

		default:
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}
	})
}
