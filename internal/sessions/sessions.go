package sessions

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID string
	DB *sql.DB // Pointer to DB for efficiency
	mu sync.Mutex
}

func New(db *sql.DB) *Session {
	return &Session{DB: db}
}

// Create a unique session ID using UUID
func UniqueID(s *Session) {
	s.ID = uuid.NewString()
}

func (s *Session) Save() error {
	s.mu.Lock() // Locking to avoid concurrent writes
	defer s.mu.Unlock()

	query := `INSERT INTO sessions (id, created_at, expires_at) VALUES (?, ?, ?)`
	_, err := s.DB.Exec(query, s.ID, time.Now(), time.Now().Add(30*time.Minute))
	if err != nil {
		return fmt.Errorf("failed to save session: %v", err)
	}

	return nil
}

func (s *Session) Send(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    s.ID,
		Path:     "/",
		HttpOnly: true, // For security
		Secure:   true, // Only send over HTTPS (ensure you're using HTTPS)
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

func (s *Session) Clean() error {
	query := `DELETE FROM sessions WHERE expires_at < ?`
	_, err := s.DB.Exec(query, time.Now())
	return err
}
