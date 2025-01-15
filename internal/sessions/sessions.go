package sessions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Token    string
	DB       *sql.DB
	mu       sync.Mutex
	Data     map[string]any // Store session data as key-value pairs
	Lifetime time.Duration
	Cookie   struct {
		Name string
	}
	ErrorFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func New() *Session {
	return &Session{
		Data: make(map[string]any),
	}
}

// Create a unique session ID using UUID
func UniqueID(s *Session) {
	s.Token = uuid.NewString()
}

// Put adds a key-value pair to the session data.
func (s *Session) Put(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Data[key] = value
}

// Save the session to the database
func (s *Session) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expiryTime := time.Now().Add(s.Lifetime)

	// Serialize session data to JSON format
	dataBytes, err := json.Marshal(s.Data)
	if err != nil {
		return fmt.Errorf("failed to serialize session data: %v", err)
	}

	query := `INSERT OR REPLACE INTO sessions (token, data, expiry) VALUES (?, ?, ?)`
	_, err = s.DB.Exec(query, s.Token, dataBytes, expiryTime)
	if err != nil {
		return fmt.Errorf("failed to save session: %v", err)
	}

	return nil
}

// Load the session from the database using the token
func (s *Session) Load() error {
	query := `SELECT data FROM sessions WHERE token = ? AND expiry > ?`
	row := s.DB.QueryRow(query, s.Token, time.Now())

	var dataBytes []byte
	err := row.Scan(&dataBytes)
	if err != nil {
		return fmt.Errorf("failed to load session: %v", err)
	}

	// Deserialize JSON data into the session map
	err = json.Unmarshal(dataBytes, &s.Data)
	if err != nil {
		return fmt.Errorf("failed to deserialize session data: %v", err)
	}

	return nil
}

// Middleware to load and save sessions
// LoadAndSave() middleware checks each
// incoming request for a session cookie. If a session cookie is present,
// it reads the session token
// and retrieves the corresponding session data from the database
// (while also checking that the
// session hasn’t expired). It then adds the session data to the request
// context so it can be used
// in your handlers.
func (s *Session) LoadAndSave(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		cookie, err := r.Cookie(s.Cookie.Name)
		if err == nil {
			token = cookie.Value
		}

		if token != "" {
			s.Token = token
			err := s.Load()
			if err != nil && s.ErrorFunc != nil {
				s.ErrorFunc(w, r, err)
				return
			}
		}

		next.ServeHTTP(w, r)

		s.Save()
		s.Send(w)
	})
}

// Send the session token to the client via a cookie
func (s *Session) Send(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     s.Cookie.Name,
		Value:    s.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

// PopString retrieves the value for the given key and deletes it from the session.
// If the key does not exist, it returns an empty string.
func (s *Session) PopString(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Retrieve and remove the key from the map
	if value, exists := s.Data[key]; exists {
		delete(s.Data, key)

		// Convert value to string and return
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}

	// Return empty string if key is missing or not a string
	return ""
}
