package handlers

import (
	"fmt"
	"myforum/internal/config"
	"net/http"
)

func NewPost(app *config.Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	w.Header().Set("Allow", http.MethodPost)
		// 	app.ClientError(w, http.StatusMethodNotAllowed)
		// 	return
		// }
		// Create some variables holding dummy data. We'll remove these later on
		// during the build.
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		// Pass the data to the SnippetModel.Insert() method, receiving the
		// ID of the new record back.
		app.InfoLog.Println("start insertion")
		id, err := app.ForumModel.Insert(title, content)
		if err != nil {
			app.ServerError(w, err)
			return
		}
		app.InfoLog.Println("end insertion")
		// Redirect the user to the relevant page for the snippet.
		//http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
	})
}
