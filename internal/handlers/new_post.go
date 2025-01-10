package handlers

import (
	"fmt"
	"html/template"
	"log"
	models "myforum/internal/models"
	"net/http"
	"time"
)

func (app *Application) NewPost(userId int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			path := "./web/templates/"
			files := []string{
				path + "base.html",
				path + "pages/creat_post.html",
			}
			tmpl, err := template.ParseFiles(files...)
			if err != nil {
				log.Printf("Error parsing template: %v", err)
				http.Error(w, "Internal Server Errorr", http.StatusInternalServerError)
				return
			}
			categories, err := GetCategories(app.ForumModel.DB, false)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "get categories", http.StatusInternalServerError)
				return
			}
			feed := struct {
				Style      string
				Categories []models.Category
			}{
				Style:      "new_post.css",
				Categories: categories,
			}
			err = tmpl.ExecuteTemplate(w, "base", feed)
			if err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				log.Printf("Error parsing form: %v", err)
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			post := &models.Post{
				Title:      r.PostFormValue("title"),
				Content:    r.PostFormValue("content"),
				Created_At: time.Now(),
				UserId:     userId,
			}
			// category := r.PostFormValue("category")
			categories := r.Form["category"]
			app.InfoLog.Println(post)
			_, err := app.ForumModel.InsertPost(*post, categories)
			if err != nil {
				http.Error(w, "internal", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		// if r.Method != http.MethodPost {
		// 	w.Header().Set("Allow", http.MethodPost)
		// 	app.ClientError(w, http.StatusMethodNotAllowed)
		// 	return
		// }
		// Create some variables holding dummy data. We'll remove these later on
		// during the build.
		// title := "O snail"
		// content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		// // Pass the data to the SnippetModel.Insert() method, receiving the
		// // ID of the new record back.
		// app.InfoLog.Println("start insertion")
		// _, err := app.ForumModel.Insert(title, content)
		// if err != nil {
		// 	app.ServerError(w, err)
		// 	return
		// }
		// app.InfoLog.Println("end insertion")
		// Redirect the user to the relevant page for the snippet.
		// http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
	})
}
