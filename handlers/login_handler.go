package handlers

import (
	"net/http"

	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

func Login(sm *middleware.SessionManager, db *middleware.AuthDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			http.NotFound(w, r)
			return
		}

		reqID := middleware.GetReqID(r)
		sess := sm.Start(w, r)

		switch r.Method {
		case "GET":
			token := csrfToken()
			sess.Set(middleware.SessionFormToken, token)

			pd := &PageData{
				Title: "Log in",
				Token: token,
			}

			tryRenderPage(w, r, PageLogin, pd)
			return

		case "POST":
			if err := r.ParseForm(); err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			sessToken := sess.GetString(middleware.SessionFormToken)
			if sessToken == "" {
				errorPage(w, http.StatusBadRequest)
				return
			}

			// Remove token from session to prevent duplicate submissions
			sess.Delete(middleware.SessionFormToken)

			// Validate form
			if err := r.ParseForm(); err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			// Ignore submissions with invalid tokens
			if formToken := r.PostFormValue("token"); sessToken != formToken {
				errorPage(w, http.StatusBadRequest)
				return
			}

			user := r.PostFormValue("username")
			plainTextPass := r.PostFormValue("password")
			if err := db.Authorize(user, plainTextPass); err != nil {
				token := csrfToken()
				sess.Set(middleware.SessionFormToken, token)

				pd := &PageData{
					Title:    "Log in",
					Token:    token,
					ErrorMsg: "Incorrect username or password.",
				}

				tryRenderPage(w, r, PageLogin, pd)
				return
			}

			// Login and return user to their previous page
			sess.Set(middleware.SessionUsername, user)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		default:
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}
	})
}
