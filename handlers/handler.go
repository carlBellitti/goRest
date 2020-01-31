package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"test.com/goRest/storage"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// A (simple) example of our application-wide configuration.
type Env struct {
	// DB   *sql.DB
	// Port string
	// Host string
	Storage *storage.StorageRepo
}

type CookieInfo struct {
	ShouldSet    bool
	ShouldExpire bool
	ID           string
	SessionID    string
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H                       func(e *Env, w http.ResponseWriter, r *http.Request) error
	Method                  string
	ShouldAuthenticateRoute bool
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != h.Method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if h.ShouldAuthenticateRoute {
		_, err := IsCookieAuthenticated(r, h.Env.Storage)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func HandlerParseResponseAndSetCookie(w http.ResponseWriter, r []byte, e error, c CookieInfo) error {
	if e != nil {
		//http.Error(w, e.Error(), http.StatusInternalServerError)
		return e
	}
	if c.ShouldSet {
		var expire time.Time
		if c.ShouldExpire {
			expire = time.Now().AddDate(0, 0, -1)
		} else {
			expire = time.Now().AddDate(1, 0, 0)
		}
		cookie := &http.Cookie{
			Name:    "sessionid",
			Value:   c.SessionID,
			Expires: expire,
		}
		http.SetCookie(w, cookie)
		cookie = &http.Cookie{
			Name:    "id",
			Value:   c.ID,
			Expires: expire,
		}
		http.SetCookie(w, cookie)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
	return nil
}

type AuthCookie = struct {
	ID        string
	sessionID string
}

func IsCookieAuthenticated(r *http.Request, s *storage.StorageRepo) (AuthCookie, error) {
	idCookie, err := r.Cookie("id")
	if err != nil {
		return AuthCookie{"", ""}, err
	}
	id := idCookie.Value

	sessionCookie, err := r.Cookie("sessionid")
	if err != nil {
		return AuthCookie{"", ""}, err
	}
	sessionID := sessionCookie.Value

	authRouteError := s.AuthenticateRoute(id, sessionID)
	if authRouteError != nil {
		return AuthCookie{"", ""}, errors.New("Unauthorized")
	}
	return AuthCookie{id, sessionID}, nil
}

func DoesIDMatchCookie(id string, r *http.Request) bool {
	idCookie, err := r.Cookie("id")
	if err != nil || idCookie.Value != id {
		return false
	}
	return true
}
