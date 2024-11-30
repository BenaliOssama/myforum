package handlers

import (
	"myforum/internal/config"
	"net/http"
	"text/template"
)

func Home(app *config.Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound(w)
			return
		}
		// Define template path
		path := "./web/templates/"
		files := []string{
			path + "base.html",
			path + "pages/posts.html",
		}

		// Parse template files
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			serverError(w, app, err) // Use the serverError() helper.
			return
		}

		// Template data
		feed := struct {
			Style string
			Posts bool
		}{
			Style: "post.css",
			Posts: false,
		}

		// Execute the template
		err = tmpl.ExecuteTemplate(w, "base", feed)
		if err != nil {
			serverError(w, app, err) // Use the serverError() helper.
			return
		}

		app.InfoLog.Println("Home handler executed successfully")
	})
}
