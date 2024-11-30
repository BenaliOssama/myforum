package handlers

import (
	"encoding/json"
	"log"
	"myforum/internal/config"
	models "myforum/internal/models"
	"net/http"
	"strconv"
)

func PostsApi(app *config.Application, isUser bool, userId int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := app.ForumModel.DB
		if r.URL.Path != "/api/posts" {
			http.Error(w, "not found", 404)
		}
		if r.Method != "GET" {
			http.Error(w, "method Not allowed", http.StatusMethodNotAllowed)
		}
		id := r.FormValue("id")
		if id != "" {
			idint, _ := strconv.Atoi(id)
			post := models.Read_Post(idint, db, isUser, userId)
			post.Categories, _ = models.GetPostCategories(post.PostId, db, userId)
			json, err := json.Marshal(post)
			if err != nil {
				log.Fatal(err)
			}
			_, _ = w.Write(json)
			return
		}
		lastindex := models.Get_Last(db)
		json, err := json.Marshal(lastindex)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = w.Write(json)
	})
}
