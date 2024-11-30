package models

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?

// This will insert a new snippet into the database.
func (m *ForumModel) InsertPost(post Post, categories []string) (int64, error) {
	transaction, err := m.DB.Begin()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting transaction:", err)
		return 0, err
	}
	stmt, err := transaction.Prepare(`INSERT INTO posts(user_id ,title,content) Values (?,?,?);`)
	if err != nil {
		transaction.Rollback()
		fmt.Fprintln(os.Stderr, "Error Adding post:", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.UserId, post.Title, post.Content)
	if err != nil {
		transaction.Rollback()
		fmt.Fprintln(os.Stderr, "Error Adding post:", err)
		return 0, err
	}
	lastPostID, err := result.LastInsertId()
	if err != nil {
		transaction.Rollback()
		fmt.Fprintln(os.Stderr, "error in assigning category to post", err)
		return 0, err
	}

	categories = append(categories, "2")
	err = LinkPostWithCategory(transaction, categories, lastPostID, post.UserId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		fmt.Fprintln(os.Stderr, "transaction aborted")
		return 0, err

	}
	return lastPostID, nil
}

func LinkPostWithCategory(transaction *sql.Tx, categories []string, postId int64, userId int) error {
	for _, category := range categories {
		stmt, err := transaction.Prepare(`INSERT INTO post_categories(user_id, post_id, category_id) VALUES(?, ?, ?);`)
		if err != nil {
			return err
		}
		defer stmt.Close()
		tmp, err := strconv.Atoi(category)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(userId, postId, tmp)
		if err != nil {
			return err
		}
	}
	return nil
}
func Read_Post(id int, db *sql.DB, isUser bool, userId int) *Post {
	query := `SELECT * FROM posts WHERE id = ?`
	row := db.QueryRow(query, id)
	Post := &Post{}
	err := row.Scan(&Post.PostId, &Post.UserId, &Post.Title, &Post.Content, &Post.LikeCount, &Post.DislikeCount, &Post.Created_At)
	if err != nil {
		fmt.Println(err)
	}

	if !isUser {
		Post.Clicked = false
		Post.DisClicked = false
	} else {
		Post.Clicked, Post.DisClicked = isLiked(db, userId, Post.PostId, "post")
	}
	Post.UserName, err = GetUserName(int(Post.UserId), db)
	if err != nil {
		fmt.Println(err)
	}
	return Post
}

func isLiked(db *sql.DB, userId int, postId int, target_type string) (bool, bool) {
	// Query to check for likes or dislikes for the given user and post.
	var query string
	if target_type == "post" {
		query = `SELECT type FROM likes WHERE user_id = ? AND post_id = ? AND target_type = ? LIMIT 1`
	} else if target_type == "comment" {
		query = `SELECT type FROM likes WHERE user_id = ? AND comment_id = ? AND target_type = ? LIMIT 1`
	}

	var reactionType string
	err := db.QueryRow(query, userId, postId, target_type).Scan(&reactionType)
	if err != nil {
		if err == sql.ErrNoRows {
			// No interaction found.
			fmt.Println("\033[31m", err, "\033[0m")
			return false, false
		}
		// Log error if needed and handle it appropriately.
		fmt.Println("Error   likes:", err)
		return false, false
	}

	// Determine the type of interaction.
	switch reactionType {
	case "like":
		return true, false
	case "dislike":
		return false, true
	default:
		return false, false
	}
}

func GetUserName(id int, db *sql.DB) (string, error) {
	var name string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func Get_Last(db *sql.DB) int {
	query := `SELECT MAX(id) FROM posts `
	row := db.QueryRow(query)
	result := 0
	_ = row.Scan(&result)
	return result
}

// // This will return a specific snippet based on its id.
// func (m *ForumModel) Get(id int) (*Post, error) {
// 	// Write the SQL statement we want to execute. Again, I've split it over two
// 	// lines for readability.
// 	stmt := `SELECT id, title, content, created FROM posts
// WHERE created < datetime('now') AND id = ?`

// 	// Use the QueryRow() method on the connection pool to execute our
// 	// SQL statement, passing in the untrusted id variable as the value for the
// 	// placeholder parameter. This returns a pointer to a sql.Row object which
// 	// holds the result from the database.
// 	row := m.DB.QueryRow(stmt, id)
// 	// Initialize a pointer to a new zeroed Snippet struct.
// 	s := &Post{}
// 	// Use row.Scan() to copy the values from each field in sql.Row to the
// 	// corresponding field in the Snippet struct. Notice that the arguments
// 	// to row.Scan are *pointers* to the place you want to copy the data into,
// 	// and the number of arguments must be exactly the same as the number of
// 	// columns returned by your statement.
// 	err := row.Scan(&s.Id, &s.Title, &s.Content, &s.Created)
// 	if err != nil {
// 		// If the query returns no rows, then row.Scan() will return a
// 		// sql.ErrNoRows error. We use the errors.Is() function check for that
// 		// error specifically, and return our own ErrNoRecord error
// 		// instead (we'll create this in a moment).
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, ErrNoRecord
// 		} else {
// 			return nil, err
// 		}
// 	}
// 	// If everything went OK then return the Snippet object.
// 	return s, nil
// }

// // This will return the 10 most recently created snippets.
// func (m *ForumModel) Latest() ([]*Post, error) {
// 	// Write the SQL statement we want to execute.
// 	stmt := `SELECT id, title, content, created FROM posts
// WHERE created < datetime('now') ORDER BY id DESC LIMIT 10`

// 	// Use the Query() method on the connection pool to execute our
// 	// SQL statement. This returns a sql.Rows resultset containing the result of
// 	// our query.
// 	rows, err := m.DB.Query(stmt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// We defer rows.Close() to ensure the sql.Rows resultset is
// 	// always properly closed before the Latest() method returns. This defer
// 	// statement should come *after* you check for an error from the Query()
// 	// method. Otherwise, if Query() returns an error, you'll get a panic
// 	// trying to close a nil resultset.
// 	defer rows.Close()
// 	// Initialize an empty slice to hold the Snippet structs.
// 	posts := []*Post{}
// 	// Use rows.Next to iterate through the rows in the resultset. This
// 	// prepares the first (and then each subsequent) row to be acted on by the
// 	// rows.Scan() method. If iteration over all the rows completes then the
// 	// resultset automatically closes itself and frees-up the underlying
// 	// database connection.
// 	for rows.Next() {
// 		// Create a pointer to a new zeroed Snippet struct.
// 		s := &Post{}
// 		// Use rows.Scan() to copy the values from each field in the row to the
// 		// new Snippet object that we created. Again, the arguments to row.Scan()
// 		// must be pointers to the place you want to copy the data into, and the
// 		// number of arguments must be exactly the same as the number of
// 		// columns returned by your statement.
// 		err = rows.Scan(&s.Id, &s.Title, &s.Content, &s.Created)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// Append it to the slice of snippets.
// 		posts = append(posts, s)
// 	}
// 	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
// 	// error that was encountered during the iteration. It's important to
// 	// call this - don't assume that a successful iteration was completed
// 	// over the whole resultset.
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	// If everything went OK then return the Snippets slice.
// 	return posts, nil
// }

// func (m *ForumModel) ExampleTransaction() error {
// 	// Calling the Begin() method on the connection pool creates a new sql.Tx
// 	// object, which represents the in-progress database transaction.
// 	tx, err := m.DB.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	// Defer a call to tx.Rollback() to ensure it is always called before the
// 	// function returns. If the transaction succeeds it will be already be
// 	// committed by the time tx.Rollback() is called, making tx.Rollback() a
// 	// no-op. Otherwise, in the event of an error, tx.Rollback() will rollback
// 	// the changes before the function returns.
// 	defer tx.Rollback()
// 	// Call Exec() on the transaction, passing in your statement and any
// 	// parameters. It's important to notice that tx.Exec() is called on the
// 	// transaction object just created, NOT the connection pool. Although we're
// 	// using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
// 	// exactly the same way.
// 	_, err = tx.Exec("INSERT INTO ...")
// 	if err != nil {
// 		return err
// 	}
// 	// Carry out another transaction in exactly the same way.
// 	_, err = tx.Exec("UPDATE ...")
// 	if err != nil {
// 		return err
// 	}
// 	// If there are no errors, the statements in the transaction can be committed
// 	// to the database with the tx.Commit() method.
// 	err = tx.Commit()
// 	return err
// }
