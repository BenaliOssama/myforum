package handlers

import (
	"database/sql"
	models "myforum/internal/models"
)

func GetCategories(db *sql.DB, withLikedAndCreated bool) ([]models.Category, error) {
	var result []models.Category
	var err error
	var rows *sql.Rows
	if withLikedAndCreated {
		rows, err = db.Query(`SELECT id, name, description FROM categories`)
	} else {
		rows, err = db.Query(`SELECT id, name, description FROM categories WHERE id != 1 AND id != 2`)
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row models.Category
		if err := rows.Scan(&row.Id, &row.Name, &row.Description); err == nil {
			result = append(result, row)
		} else {
			return nil, err
		}
	}
	return result, nil
}
