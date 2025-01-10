package handlers

import (
	"log"
	"myforum/internal/models"
	"net/http"
)

type Application struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	ForumModel *models.ForumModel
}

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()


	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	//mux.HandleFunc("/snippet/view", app.newPost)
	//mux.HandleFunc("/snippet/create", app.snippetCreate)
	return mux
}


/*
	fileServer := http.FileServer(http.Dir("./web/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	mux.Handle("/", handlers.Home(app))
	mux.Handle("/new_post", handlers.NewPost(app, 1))
*