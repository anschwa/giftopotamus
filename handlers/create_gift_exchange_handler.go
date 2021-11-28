package handlers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/anschwa/giftopotamus/giftex"
	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

func CreateGiftExchange(sm *middleware.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r)

		switch r.Method {
		case "GET":
			http.Redirect(w, r, "/", http.StatusFound)
			return

		case "POST":
			sess := sm.Start(w, r)
			username := sess.GetString(middleware.SessionUsername)

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

			// Get form values
			if err := r.ParseForm(); err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			var tableRows []GiftexTableRow
			tableJSON := r.PostFormValue("participants")
			if err := json.Unmarshal([]byte(tableJSON), &tableRows); err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			if len(tableRows) == 0 {
				sess.Set(middleware.SessionErrorMsg, "Oops! Your gift exchange is empty. Please add some participants and try again.")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			db, err := tableRowsToGiftExchangeDB(tableRows)
			if err := json.Unmarshal([]byte(tableJSON), &tableRows); err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			ge, err := giftex.NewGiftExchange(db.Participants, &giftex.GiftExchangeOptions{MaxPrevious: 2})
			if err != nil {
				if errors.Is(err, giftex.ErrNoSolution) {
					sess.Set(middleware.SessionTableRows, tableRows)
					sess.Set(middleware.SessionErrorMsg, "Oops! An assignment isn't possible with your current gift exchange. Please adjust your restrictions and try again.")
					http.Redirect(w, r, "/", http.StatusFound)
					return
				}

				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			resultsTable, resultsCSV, err := giftExchangeToTableRows(db, ge)
			if err != nil {
				logger.Error(reqID, err)
				errorPage(w, http.StatusInternalServerError)
				return
			}

			// Save CSV on session for download
			sess.Set(middleware.SessionResultsCSV, resultsCSV)

			// Display results
			pd := &PageData{
				Title:      "Giftopotamus.com",
				Username:   username,
				SuccessMsg: "Gift exchange created!",

				TableRows:  resultsTable,
				ResultsCSV: resultsCSV,
			}

			tryRenderPage(w, r, PageResults, pd)
			return

		default:
			errorPage(w, http.StatusMethodNotAllowed)
			return
		}
	})
}

func tableRowsToCSV(rows []GiftexTableRow) ([]byte, error) {
	// Construct CSV from rows
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Write([]string{"name", "email", "restrictions", "previous", "participating", "has"})
	for _, p := range rows {
		w.Write([]string{
			strings.TrimSpace(p.Name),
			strings.TrimSpace(p.Email),
			strings.TrimSpace(p.Restrictions),
			strings.TrimSpace(p.Previous),
			"yes", // Everyone is participating
			"",    // The Has column is required for writing out the results later
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, fmt.Errorf("Error converting []GiftexTableRow to CSV: %w", err)
	}

	return buf.Bytes(), nil
}

func tableRowsToGiftExchangeDB(rows []GiftexTableRow) (*giftex.GiftExchangeDB, error) {
	rowsCSV, err := tableRowsToCSV(rows)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(rowsCSV)
	return giftex.ReadCSV(buf)
}

func giftExchangeToTableRows(db *giftex.GiftExchangeDB, ge *giftex.GiftExchange) (resultsTable []GiftexTableRow, resultsCSV []byte, err error) {
	var resultsBuf, tmpBuf bytes.Buffer
	if err := db.WriteCSV(&resultsBuf, ge.Assignment); err != nil {
		return nil, nil, err
	}

	tmpBuf.Write(resultsBuf.Bytes()) // Get copy of results
	tmpDB, err := giftex.ReadCSV(&tmpBuf)
	if err != nil {
		return nil, nil, err
	}

	for _, p := range tmpDB.Participants {
		restrictions := make([]string, 0, len(p.Restrictions))
		for _, pid := range p.Restrictions {
			restrictions = append(restrictions, tmpDB.Participants[pid].Name)
		}

		previous := make([]string, 0, len(p.Previous))
		for _, pid := range p.Previous {
			previous = append(previous, tmpDB.Participants[pid].Name)
		}

		resultsTable = append(resultsTable, GiftexTableRow{
			Name:         p.Name,
			Email:        p.Email,
			Restrictions: strings.Join(restrictions, ", "),
			Previous:     strings.Join(previous, ", "),
			Has:          tmpDB.Participants[ge.Assignment[p.ID]].Name,
		})
	}

	return resultsTable, resultsBuf.Bytes(), nil
}
