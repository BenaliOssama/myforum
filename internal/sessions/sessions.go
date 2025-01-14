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
	Token    string
	DB       *sql.DB // Pointer to DB for efficiency
	mu       sync.Mutex
	Data     []byte        // Store session data as a byte slice (BLOB)
	Lifetime time.Duration // Lifetime of the session
	Cookie   struct {
		Name string
	}
	// Handle errors in middleware
	ErrorFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func New() *Session {
	return &Session{}
}

// Create a unique session ID using UUID
func UniqueID(s *Session) {
	s.Token = uuid.NewString() // Generate a new UUID and assign it to Token
}

// Save the session to the database
func (s *Session) Save() error {
	s.mu.Lock() // Locking to avoid concurrent writes
	defer s.mu.Unlock()

	// Calculate session expiry time based on the Lifetime variable
	expiryTime := time.Now().Add(s.Lifetime)

	// Inserting the session data into the database
	query := `INSERT INTO sessions (token, data, expiry) VALUES (?, ?, ?)`
	_, err := s.DB.Exec(query, s.Token, s.Data, expiryTime) // Use dynamic expiry time
	if err != nil {
		return fmt.Errorf("failed to save session: %v", err)
	}

	return nil
}

// Send the session token to the client via a cookie
func (s *Session) Send(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    s.Token,
		Path:     "/",
		HttpOnly: true, // For security
		Secure:   true, // Only send over HTTPS (ensure you're using HTTPS)
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

// Clean expired sessions from the database
func (s *Session) Clean() error {
	query := `DELETE FROM sessions WHERE expiry < ?`
	_, err := s.DB.Exec(query, time.Now())
	return err
}

// LoadAndSave middleware implementation (without context part)
func (s *Session) LoadAndSave(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract session token from the cookie
		var token string
		cookie, err := r.Cookie(s.Cookie.Name)
		if err == nil {
			token = cookie.Value // If cookie exists, get its value
		}

		// If session token exists, load the session from the database
		if token != "" {
			s.Token = token
			err := s.Load()
			if err != nil {
				// If there's an error loading the session, handle it
				if s.ErrorFunc != nil {
					s.ErrorFunc(w, r, err)
				}
				return
			}
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// After handler is done, save the session and send the cookie back to the client
		s.Save()
		s.Send(w)
	})
}

// Load the session from the database using the token
func (s *Session) Load() error {
	// Query to load the session data from the database
	query := `SELECT data FROM sessions WHERE token = ? AND expiry > ?`
	row := s.DB.QueryRow(query, s.Token, time.Now())

	// Retrieve the session data from the database
	err := row.Scan(&s.Data)
	if err != nil {
		return fmt.Errorf("failed to load session: %v", err)
	}

	return nil
}
