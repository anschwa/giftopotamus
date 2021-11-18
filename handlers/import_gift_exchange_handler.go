package handlers

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/anschwa/giftopotamus/giftex"
	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

const maxFileSize = 1 << 20 // 1 MiB

func ImportGiftExchange(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r)

		if r.Method != "POST" {
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}

		sess := sm.Start(w, r)
		sessToken := sess.GetString(middleware.SessionFormToken)
		if sessToken == "" {
			errorPage(w, http.StatusBadRequest)
			return
		}

		// Remove token from session to prevent duplicate submissions
		sess.Delete(middleware.SessionFormToken)

		// Ignore submissions with invalid tokens
		if formToken := r.PostFormValue("token"); sessToken != formToken {
			errorPage(w, http.StatusBadRequest)
			return
		}

		if err := r.ParseMultipartForm(maxFileSize); err != nil {
			logger.Error(reqID, err)
			errorPage(w, http.StatusInternalServerError)
			return
		}

		file, header, err := r.FormFile("csv")
		if err != nil {
			logger.Error(reqID, err)
			errorPage(w, http.StatusInternalServerError)
			return
		}
		defer file.Close()

		tableRows, err := csvToRows(file)
		if err != nil {
			logger.Error(reqID, err)
			errorPage(w, http.StatusInternalServerError)
			return
		}

		// Update session data
		sess.Set(middleware.SessionFormToken, csrfToken())
		sess.Set(middleware.SessionSuccessMsg, fmt.Sprintf("Imported %s", header.Filename))
		sess.Set(middleware.SessionTableRows, tableRows)

		http.Redirect(w, r, "/", http.StatusFound)
	})
}

type GiftexTableRow struct {
	Name         string
	Email        string
	Restrictions string
	Previous     string
	Has          string
}

func csvToRows(r io.Reader) ([]GiftexTableRow, error) {
	db, err := giftex.ReadCSV(r)
	if err != nil {
		return nil, err
	}

	rows := make([]GiftexTableRow, 0, len(db.Participants))
	for _, p := range db.Participants {
		restrictions := make([]string, 0, len(p.Restrictions))
		for _, pid := range p.Restrictions {
			restrictions = append(restrictions, db.Participants[pid].Name)
		}

		previous := make([]string, 0, len(p.Previous))
		for _, pid := range p.Previous {
			previous = append(previous, db.Participants[pid].Name)
		}

		row := GiftexTableRow{
			Name:         p.Name,
			Email:        p.Email,
			Restrictions: strings.Join(restrictions, ", "),
			Previous:     strings.Join(previous, ", "),
		}

		rows = append(rows, row)
	}

	// Sort by name
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Name < rows[j].Name
	})

	return rows, nil
}
