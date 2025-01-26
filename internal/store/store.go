package store

import (
	"database/sql"
	"time"
)

// scs.Store defines the interface for custom session stores. Any object that implements this interface can be set as the store when configuring the session.

// type Store interface {
// 	// Delete should remove the session token and corresponding data from the
// 	// session store. If the token does not exist then Delete should be a no-op
// 	// and return nil (not an error).
// 	Delete(token string) (err error)

// 	// Find should return the data for a session token from the store. If the
// 	// session token is not found or is expired, the found return value should
// 	// be false (and the err return value should be nil). Similarly, tampered
// 	// or malformed tokens should result in a found return value of false and a
// 	// nil err value. The err return value should be used for system errors only.
// 	Find(token string) (b []byte, found bool, err error)

// 	// Commit should add the session token and data to the store, with the given
// 	// expiry time. If the session token already exists, then the data and
// 	// expiry time should be overwritten.
// 	Commit(token string, b []byte, expiry time.Time) (err error)
// }

// type IterableStore interface {
// 	// All should return a map containing data for all active sessions (i.e.
// 	// sessions which have not expired). The map key should be the session
// 	// token and the map value should be the session data. If no active
// 	// sessions exist this should return an empty (not nil) map.
// 	All() (map[string][]byte, error)
// }

// SessionStore is a custom session store that implements the Store and IterableStore interfaces
type SessionStore struct {
	db *sql.DB
}

// NewSessionStore initializes and returns a new SessionStore instance
func New(db *sql.DB) *SessionStore {
	return &SessionStore{db: db}
}

// Delete removes the session with the given token from the store
func (s *SessionStore) Delete(token string) error {
	_, err := s.db.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

// Find retrieves the session data for the given token
func (s *SessionStore) Find(token string) ([]byte, bool, error) {
	var data []byte
	var expiry time.Time

	err := s.db.QueryRow("SELECT data, expiry FROM sessions WHERE token = ?", token).Scan(&data, &expiry)
	if err == sql.ErrNoRows {
		return nil, false, nil // Token not found
	}
	if err != nil {
		return nil, false, err // Other errors
	}

	if time.Now().After(expiry) {
		// Token is expired
		return nil, false, nil
	}

	return data, true, nil
}

// Commit adds or updates the session data for the given token with expiry
func (s *SessionStore) Commit(token string, data []byte, expiry time.Time) error {
	_, err := s.db.Exec(`
        INSERT OR REPLACE INTO sessions (token, data, expiry)
        VALUES (?, ?, ?)
    `, token, data, expiry)
	return err
}
