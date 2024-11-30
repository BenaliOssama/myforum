package modules

import (
	"database/sql"
	"log"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type ForumModel struct {
	DB *sql.DB
}

// for a given DSN.
func OnpenDB(dsn string, infoLog *log.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	infoLog.Println("data base opened/created seccusfuly")
	if err = db.Ping(); err != nil {
		return nil, err
	}
	infoLog.Println("data base connected seccusfuly")
	err = createTables(db)
	if err != nil {
		return db, err
	}
	infoLog.Println("All tables Created succesfully")
	return db, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(PostsTable)
	if err != nil {
		return err
	}
	return nil
}
