package main

import (
	"net/http"
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
// The routes() method returns a servemux containing our application routes.
// func (app *application) routes() *http.ServeMux {
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// pass home through the middleware of sessions
	home := app.sessionManager.LoadAndSave(http.HandlerFunc(app.home))
	mux.Handle("/", home)

	// pass through the middleware of sessions
	snippetView := app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetView))
	mux.Handle("/snippet/view", snippetView)

	mux.HandleFunc("/snippet/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// pass through the middleware of sessions
			snippetCreate := app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetCreate))
			snippetCreate.ServeHTTP(w, r)
		case http.MethodPost:
			// pass through the middleware of sessions
			snippetCreatePost := app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetCreate))
			snippetCreatePost.ServeHTTP(w, r)
		default:
			app.clientError(w, http.StatusMethodNotAllowed)
		}
	})

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware.
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.

	// Wrap the existing chain with the logRequest middleware.
	//return app.logRequest(secureHeaders(mux))
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
