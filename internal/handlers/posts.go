package handlers

// func Posts(app *config.Application) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 		if err != nil || id < 1 {
// 			app.NotFound(w)
// 			return
// 		}
// 		// Use the SnippetModel object's Get method to retrieve the data for a
// 		// specific record based on its ID. If no matching record is found,
// 		// return a 404 Not Found response.
// 		snippet, err := app.ForumModel.Get(id)
// 		if err != nil {
// 			if errors.Is(err, models.ErrNoRecord) {
// 				app.NotFound(w)
// 			} else {
// 				app.ServerError(w, err)
// 			}
// 			return
// 		}
// 		// Write the snippet data as a plain-text HTTP response body.
// 		fmt.Fprintf(w, "%+v", snippet)
// 	})
// }
