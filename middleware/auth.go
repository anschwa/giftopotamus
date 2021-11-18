package middleware

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthDBUserNotFound = errors.New("Error: user not found in AuthDB")
)

type AuthDB struct {
	mu sync.Mutex

	entries map[string]string
}

func NewAuthDB(entries map[string]string) *AuthDB {
	return &AuthDB{
		entries: entries,
	}
}

func (db *AuthDB) Authorize(user, pass string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	hash, ok := db.entries[user]
	if !ok {
		return ErrAuthDBUserNotFound
	}

	return compare(hash, pass)
}

func (db *AuthDB) SetPassword(user, pass string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	hash, err := hash(pass)
	if err != nil {
		return err
	}

	db.entries[user] = hash
	return nil
}

func compare(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

func hash(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}
