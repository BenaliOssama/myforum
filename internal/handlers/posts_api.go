package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) PostsApi(isUser bool, userId int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/posts" {
			http.Error(w, "not found", 404)
		}
		if r.Method != "GET" {
			http.Error(w, "method Not allowed", http.StatusMethodNotAllowed)
		}
		id := r.FormValue("id")
		if id != "" {
			idint, _ := strconv.Atoi(id)
			post := app.ForumModel.Read_Post(idint)
			json, err := json.Marshal(post)
			if err != nil {
				log.Fatal(err)
			}
			_, _ = w.Write(json)
			return
		}
	})
}
