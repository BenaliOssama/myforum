package main

import (
	"flag"
	"log"
	config "myforum/internal/config"
	handlers "myforum/internal/handlers"
	"net/http"
	"os"
)

// Initialize a new instance of our application struct, containing the
// dependencies.

func main() {
	//###################### config ##############################//
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	flag.Parse()

	app := &config.Application{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	mux := http.NewServeMux()

	//######################## Sever ##############################//
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/", handlers.Home(app))

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLog,
		Handler:  mux,
	}

	app.InfoLog.Printf("Starting server on http://%s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}
