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


func (m *ForumModel) Read_Post(id int) *Post {
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
