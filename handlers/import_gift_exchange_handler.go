package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

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

// csvToRows reads all rows in a given CSV for displaying as a table.
// We want to preserve all the data given by the user without making
// assumptions about the validity. For example, giftex.ReadCSV
// incorrectly ignore column data that include participants who have
// not been entered into the table yet.
func csvToRows(r io.Reader) ([]GiftexTableRow, error) {
	csvReader := csv.NewReader(r)
	csvReader.FieldsPerRecord = -1 // Allow empty columns
	csvReader.Comma = ','

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading csv: %w", err)
	}

	numRecords := len(records)
	if numRecords < 2 {
		return nil, fmt.Errorf("csv must include headers and at least one entry")
	}

	trimLower := func(s string) string { return strings.TrimSpace(strings.ToLower(s)) }

	// The first record contains the column headers
	cols := make(map[string]int, len(records[0]))
	for i, v := range records[0] {
		cols[trimLower(v)] = i
	}

	tableRows := make([]GiftexTableRow, 0, numRecords)
	for i := 0; i < numRecords; i++ {
		row := records[i]

		// Skip non-participants
		if participating := trimLower(row[cols["participating"]]) == "yes"; !participating {
			continue
		}

		getCol := func(key string) string {
			return strings.TrimSpace(row[cols[key]])
		}

		tr := GiftexTableRow{
			Name:         getCol("name"),
			Email:        getCol("email"),
			Restrictions: getCol("restrictions"),
			Previous:     getCol("previous"),
		}

		tableRows = append(tableRows, tr)
	}

	// Sort rows by name
	sort.Slice(tableRows, func(i, j int) bool {
		return tableRows[i].Name < tableRows[j].Name
	})

	return tableRows, nil
}
