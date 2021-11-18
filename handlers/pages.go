package handlers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

var isDev bool

func init() {
	isDev = os.Getenv("APP_ENV") != "production"
}

const (
	PageLogin   = "login"
	PageLogout  = "logout"
	PageGiftex  = "giftex"
	PageResults = "results"
)

var templates = map[string]*template.Template{
	PageLogin:   parsePage(PageLogin),
	PageLogout:  parsePage(PageLogout),
	PageGiftex:  parsePage(PageGiftex),
	PageResults: parsePage(PageResults),
}

type PageData struct {
	Title                string
	Username             string
	Token                string
	ErrorMsg, SuccessMsg string

	TableRows  []GiftexTableRow
	ResultsCSV []byte
}

func parseTemplates(pages ...string) *template.Template {
	tmpl := make([]string, len(pages))
	for i, p := range pages {
		tmpl[i] = "templates/" + p + ".html.tmpl" // keep it simple
	}

	return template.Must(template.ParseFiles(tmpl...))
}

func parsePage(page string) *template.Template {
	return parseTemplates(page, "meta", "header", "footer")
}

func renderPage(w http.ResponseWriter, page string, d *PageData) error {
	if isDev { // Re-parse templates for each request during development only
		templates[page] = parsePage(page)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	return templates[page].ExecuteTemplate(w, page, d)
}

// tryRenderPage renders page and logs error on failure
func tryRenderPage(w http.ResponseWriter, r *http.Request, page string, d *PageData) {
	reqID := middleware.GetReqID(r)

	if err := renderPage(w, page, d); err != nil {
		logger.Error(reqID, err)

		c := http.StatusInternalServerError
		http.Error(w, http.StatusText(c), c)
		return
	}
}

func errorPage(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func csrfToken() string {
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
