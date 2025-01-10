package handlers

import (
	"html/template"
	"net/http"
)

func (app *Application) Home() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			app.NotFound(w)
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
			app.ServerError(w, err) // Use the serverError() helper.
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
			app.ServerError(w, err) // Use the serverError() helper.
			return
		}

		app.InfoLog.Println("Home handler executed successfully")
	})
}
