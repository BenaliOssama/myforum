package main

import (
	"myforum/ui"
	"net/http"
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
// The routes() method returns a servemux containing our application routes.
// func (app *application) routes() *http.ServeMux {
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Take the ui.Files embedded filesystem and convert it to a http.FS type so
	// that it satisfies the http.FileSystem interface. We then pass that to the
	// http.FileServer() function to create the file server handler.
	fileServer := http.FileServer(http.FS(ui.Files))
	// Our static files are contained in the "static" folder of the ui.Files
	// embedded filesystem. So, for example, our CSS stylesheet is located at
	// "static/css/main.css". This means that we now longer need to strip the
	// prefix from the request URL -- any requests that start with /static/ can
	// just be passed directly to the file server and the corresponding static
	// file will be served (so long as it exists).
	mux.Handle("/static/", fileServer)

	// Add a new GET /ping route.
	mux.Handle("/ping", http.HandlerFunc(ping))

	mux.Handle("/", app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.home))))
	mux.Handle("/snippet/view", app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.snippetView))))
	mux.Handle("/user/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// pass through the middleware of sessions
			userSignup := app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.userSignup)))
			userSignup.ServeHTTP(w, r)
		case http.MethodPost:
			// pass through the middleware of sessions
			userSignUpPost := app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.userSignupPost)))
			userSignUpPost.ServeHTTP(w, r)
		default:
			app.clientError(w, http.StatusMethodNotAllowed)
		}
	}))
	mux.Handle("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// pass through the middleware of sessions
			userLogin := app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.userLogin)))
			userLogin.ServeHTTP(w, r)
		case http.MethodPost:
			// pass through the middleware of sessions
			userLoginPost := app.sessionManager.LoadAndSave(app.authenticate(http.HandlerFunc(app.userLoginPost)))
			userLoginPost.ServeHTTP(w, r)
		default:
			app.clientError(w, http.StatusMethodNotAllowed)
		}
	}))
	/*______________________________________authentication first___________________________*/
	mux.Handle("/snippet/create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// pass through the middleware of sessions
			snippetCreate := app.sessionManager.LoadAndSave(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))
			snippetCreate.ServeHTTP(w, r)
		case http.MethodPost:
			// pass through the middleware of sessions
			snippetCreatePost := app.sessionManager.LoadAndSave(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost))))
			snippetCreatePost.ServeHTTP(w, r)
		default:
			app.clientError(w, http.StatusMethodNotAllowed)
		}
	}))
	// Add the five new routes, all of which use our 'dynamic' middleware chain.
	mux.Handle("/user/logout", app.sessionManager.LoadAndSave(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost)))))
	// Pass the servemux as the 'next' parameter to the secureHeaders middleware.
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.

	// Wrap the existing chain with the logRequest middleware.
	//return app.logRequest(secureHeaders(mux))
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
