package handlers

import (
	"fmt"
	"myforum/internal/config"
	"net/http"
	"runtime/debug"
)

func serverError(w http.ResponseWriter, app *config.Application, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//errLog.Println(trace)
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func notFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}
