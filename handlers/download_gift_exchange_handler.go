package handlers

import (
	"net/http"

	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

func DownloadGiftExchange(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}

		reqID := middleware.GetReqID(r)
		sess := sm.Start(w, r)

		// Get results from session
		resultsCSV, err := sess.Get(middleware.SessionResultsCSV)
		if err != nil {
			errorPage(w, http.StatusBadRequest)
			return
		}

		file, ok := resultsCSV.([]byte)
		if !ok {
			logger.Error(reqID, err)
			errorPage(w, http.StatusInternalServerError)
			return
		}

		// Write file to client
		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		w.Write(file)
	})
}
