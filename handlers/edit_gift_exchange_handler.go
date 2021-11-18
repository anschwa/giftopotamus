package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

func EditGiftExchange(sm *middleware.SessionManager) http.Handler {
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

		// Get form values
		if err := r.ParseForm(); err != nil {
			logger.Error(reqID, err)
			errorPage(w, http.StatusInternalServerError)
			return
		}

		// Name is required
		participantName := r.PostFormValue("name")
		if len(strings.TrimSpace(participantName)) == 0 {
			sess.Set(middleware.SessionErrorMsg, "Name is required")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		var tableRows []GiftexTableRow
		if v, _ := sess.Get(middleware.SessionTableRows); v != nil {
			tableRows = v.([]GiftexTableRow)
		}

		nameMap := make(map[string]int, len(tableRows))
		for i, v := range tableRows {
			nameMap[v.Name] = i
		}

		// Check if 'edit' or 'remove'
		switch v := r.PostFormValue("action"); v {
		case "edit":
			// Names must be unique
			var participantIdx *int
			if idx, err := strconv.Atoi(r.PostFormValue("index")); err == nil {
				participantIdx = &idx
			}

			// New rows shouldn't have an existing index and existing rows should match the provided index
			if idx, ok := nameMap[participantName]; (participantIdx == nil && ok) || (participantIdx != nil && *participantIdx != idx) {
				sess.Set(middleware.SessionErrorMsg, fmt.Sprintf("%s is already taken", participantName))
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			row := GiftexTableRow{
				Name:         participantName,
				Email:        r.PostFormValue("email"),
				Restrictions: r.PostFormValue("restrictions"),
			}

			// Update existing participant or insert new row into table
			var msg string
			if participantIdx != nil {
				tableRows[*participantIdx] = row
				msg = fmt.Sprintf("Updated %s's info.", participantName)
			} else {
				tableRows = append(tableRows, row)
				msg = fmt.Sprintf("Added %s to the gift exchange!", participantName)
			}

			sess.Set(middleware.SessionSuccessMsg, msg)

		case "remove":
			// Ignore rows that don't exist
			idx, ok := nameMap[participantName]
			if !ok {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			tableRows = append(tableRows[:idx], tableRows[idx+1:]...)

			msg := fmt.Sprintf("Removed %s from the gift exchange.", participantName)
			sess.Set(middleware.SessionSuccessMsg, msg)

		default:
			logger.Error(reqID, fmt.Errorf("Error: expected action to be edit or remove; got: %q", v))
			errorPage(w, http.StatusInternalServerError)
			return
		}

		// Sort rows by name
		sort.Slice(tableRows, func(i, j int) bool {
			return tableRows[i].Name < tableRows[j].Name
		})

		// Update the session with new table data
		sess.Set(middleware.SessionTableRows, tableRows)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
