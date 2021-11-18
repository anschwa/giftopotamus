package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	SessionUsername   = "username"
	SessionFormToken  = "form_token"
	SessionSuccessMsg = "success_msg"
	SessionErrorMsg   = "error_msg"
	SessionTableRows  = "table_rows"
	SessionResultsCSV = "results_csv"
)

// SessionManager manages all active sessions on the web server.
type SessionManager struct {
	name       string
	mu         sync.Mutex
	ttlSeconds int // 0 will expire the session when the browser tab is closed

	sessions map[string]*Session
}

func NewSessionManager(cookieName string, ttlSeconds int) *SessionManager {
	return &SessionManager{
		name:       cookieName,
		mu:         sync.Mutex{},
		ttlSeconds: ttlSeconds,
		sessions:   make(map[string]*Session),
	}
}

// Start creates a new session with request r.
func (m *SessionManager) Start(w http.ResponseWriter, r *http.Request) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	newSess := func() *Session {
		sess := NewSession()
		c := http.Cookie{
			Name:     m.name,
			Value:    sess.id,
			Path:     "/",
			MaxAge:   m.ttlSeconds,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}

		http.SetCookie(w, &c)
		m.sessions[sess.id] = sess

		return sess
	}

	cookie, err := r.Cookie(m.name)

	// Create a new session if none exist
	if err != nil || cookie.Value == "" {
		return newSess()
	}

	sid := cookie.Value
	sess, ok := m.sessions[sid]

	// Return existing session
	if ok {
		return sess
	}

	// If no active session is found, we need to overwrite the
	// existing cookie from the browser with a new session
	return newSess()
}

// End destroys the session associated with request r
func (m *SessionManager) End(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.name)
	if err != nil || cookie.Value == "" {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, cookie.Value)
	c := http.Cookie{
		Name:     m.name,
		Path:     "/",
		MaxAge:   -1, // expire immediately
		Expires:  time.Now(),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &c)
}

type Session struct {
	mu    sync.Mutex
	id    string
	items map[string]interface{}
}

func NewSession() *Session {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}

	return &Session{
		id:    base64.URLEncoding.EncodeToString(b),
		items: make(map[string]interface{}),
	}
}

var ErrSessionItemNotFound = errors.New("Error: item not found in session")

func (s *Session) Get(key string) (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.items[key]
	if !ok {
		return nil, ErrSessionItemNotFound
	}

	return value, nil
}

// GetString returns zero value if key is not found
func (s *Session) GetString(key string) string {
	var value string

	// Make sure value exists and is not empty interface{}
	if v, err := s.Get(key); err == nil && v != nil {
		if vv, ok := v.(string); ok {
			value = vv
		}
	}

	return value
}

func (s *Session) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = value
}

func (s *Session) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.items, key)
}
