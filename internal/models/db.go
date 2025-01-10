package models

import (
	"database/sql"
	"errors"
	"log"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type ForumModel struct {
	DB *sql.DB
}

// ## errors
var ErrNoRecord = errors.New("models: no matching record found")

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
	// Calling the Begin() method on the connection pool creates a new sql.Tx
	// object, which represents the in-progress database transaction.
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// Defer a call to tx.Rollback() to ensure it is always called before the
	// function returns. If the transaction succeeds it will be already be
	// committed by the time tx.Rollback() is called, making tx.Rollback() a
	// no-op. Otherwise, in the event of an error, tx.Rollback() will rollback
	// the changes before the function returns.
	defer tx.Rollback()
	// Call Exec() on the transaction, passing in your statement and any
	// parameters. It's important to notice that tx.Exec() is called on the
	// transaction object just created, NOT the connection pool. Although we're
	// using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
	// exactly the same way.
	_, err = tx.Exec(PostsTable)
	if err != nil {
		return err
	}
	// If there are no errors, the statements in the transaction can be committed
	// to the database with the tx.Commit() method.
	err = tx.Commit()
	return err
}
