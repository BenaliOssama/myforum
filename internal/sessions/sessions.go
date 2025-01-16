package scs

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Store is the interface for session stores.
type Store interface {
	Delete(token string) error
	Find(token string) ([]byte, bool, error)
	Commit(token string, data []byte, expiry time.Time) error
}

// SessionManager holds the configuration settings for session management.
type SessionManager struct {
	Lifetime time.Duration
	Store    Store
	Cookie   SessionCookie
}

// SessionCookie contains the configuration settings for session cookies.
type SessionCookie struct {
	Name   string
	Path   string
	Secure bool
}

// New creates a new SessionManager instance.
func New() *SessionManager {
	return &SessionManager{
		Lifetime: 24 * time.Hour,
		Cookie: SessionCookie{
			Name:   "session",
			Path:   "/",
			Secure: false,
		},
	}
}

// LoadAndSave provides middleware to load and save session data.
func (s *SessionManager) LoadAndSave(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in LoadAndSave")
		ctx, err := s.loadSession(r.Context(), r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

		fmt.Println("out LoadAndSave")
		s.saveSession(w, r)
	})
}

// loadSession retrieves session data from the store and adds it to the context.
func (s *SessionManager) loadSession(ctx context.Context, r *http.Request) (context.Context, error) {
	token, err := s.getSessionToken(r)
	if err != nil {
		return nil, err
	}

	data, found, err := s.Store.Find(token)
	if err != nil {
		return nil, err
	}
	if !found {
		token = uuid.NewString()
		data, _ = s.encodeSessionData(map[string]interface{}{})
	}

	fmt.Println("in load session")
	sessionData, _ := s.decodeSessionData(data)
	fmt.Println("out load session")

	ctx = context.WithValue(ctx, "session", sessionData)
	ctx = context.WithValue(ctx, "token", token)
	return ctx, nil

	//return context.WithValue(ctx, "session", sessionData).WithValue("token", token), nil
	//return context.WithValue(ctx, "session", sessionData), nil
}

// saveSession saves the session and writes the session cookie to the response.
func (s *SessionManager) saveSession(w http.ResponseWriter, r *http.Request) {
	sessionData := r.Context().Value("session").(map[string]interface{})

	token := r.Context().Value("token").(string)
	data, _ := s.encodeSessionData(sessionData)

	expiry := time.Now().Add(s.Lifetime)
	err := s.Store.Commit(token, data, expiry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:    s.Cookie.Name,
		Value:   token,
		Path:    s.Cookie.Path,
		Expires: expiry,
		Secure:  s.Cookie.Secure,
	}

	http.SetCookie(w, cookie)
}

// getSessionToken retrieves the session token from the request cookie.
func (s *SessionManager) getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(s.Cookie.Name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// encodeSessionData encodes session data to binary.
func (s *SessionManager) encodeSessionData(data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(data)
	return buf.Bytes(), err
}

// decodeSessionData decodes binary data to session data.
func (s *SessionManager) decodeSessionData(data []byte) (map[string]interface{}, error) {
	var sessionData map[string]interface{}
	buf := bytes.NewReader(data)
	err := gob.NewDecoder(buf).Decode(&sessionData)
	return sessionData, err
}

// Put adds a key-value pair to the session data.
func (s *SessionManager) Put(ctx context.Context, key string, value interface{}) context.Context {
	sessionData := ctx.Value("session").(map[string]interface{})
	sessionData[key] = value

	return ctx
}

// PopString removes a key-value pair and returns the value as a string.
func (s *SessionManager) PopString(ctx context.Context, key string) string {
	sessionData := ctx.Value("session").(map[string]interface{})
	value, ok := sessionData[key]
	if !ok {
		return ""
	}

	delete(sessionData, key)

	return value.(string)
}
