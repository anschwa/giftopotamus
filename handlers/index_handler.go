package handlers

import (
	"net/http"

	"github.com/anschwa/giftopotamus/middleware"
)

func Index(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method != "GET" {
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}

		sess := sm.Start(w, r)
		username := sess.GetString(middleware.SessionUsername)
		sucMsg := sess.GetString(middleware.SessionSuccessMsg)
		errMsg := sess.GetString(middleware.SessionErrorMsg)

		var rows []GiftexTableRow
		if v, err := sess.Get(middleware.SessionTableRows); err == nil && v != nil {
			if vv, ok := v.([]GiftexTableRow); ok {
				rows = vv
			}
		}

		// Set csrf token
		token := csrfToken()
		sess.Set(middleware.SessionFormToken, token)

		pd := &PageData{
			Title:      "Giftopotamus.com",
			Username:   username,
			Token:      token,
			SuccessMsg: sucMsg,
			ErrorMsg:   errMsg,
			TableRows:  rows,
		}

		tryRenderPage(w, r, PageGiftex, pd)

		// Clear status after showing once
		sess.Delete(middleware.SessionSuccessMsg)
		sess.Delete(middleware.SessionErrorMsg)
	})
}
