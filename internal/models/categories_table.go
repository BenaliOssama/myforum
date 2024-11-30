package models

import (
	"database/sql"
	"fmt"
)

func GetPostCategories(idPost int, db *sql.DB, userId int) ([]string, error) {
	var categories []string
	stmt, err := db.Prepare(`SELECT categories.name 
	FROM post_categories
	JOIN categories ON categories.id = post_categories.category_id
	WHERE (post_categories.category_id = 1 OR post_categories.category_id = 2) 
	AND post_categories.user_id = ? 
	AND post_categories.post_id = ?
	UNION
	SELECT categories.name 
	FROM post_categories
	JOIN categories ON categories.id = post_categories.category_id
	WHERE post_categories.category_id != 1 
	AND post_categories.category_id != 2 
	AND post_categories.post_id = ?;
`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, idPost, idPost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
