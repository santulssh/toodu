package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

var (
	ErrNotFoundSource      = errors.New("Failed to open SQLite database.")
	ErrConnectionFailed    = errors.New("Failed to connect to SQLite database.")
	ErrFailedToCreateTable = errors.New("Failed to create a table.")
)

func NewSQLiteConnection(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)")
	if err != nil {
		return nil, fmt.Errorf(ErrNotFoundSource.Error()+":\n\n%w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf(ErrConnectionFailed.Error()+":\n\n%w", err)
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

const (
	queryCreateTasksTable = `
	CREATE TABLE IF NOT EXISTS tasks (
				id TEXT PRIMARY KEY,
				title TEXT NOT NULL,
				description TEXT,
				priority TEXT NO NULL,
				status TEXT NO NULL,
				created_at DATETIME NOT NULL,
      	modified_at DATETIME NOT NULL
		);
	`
)

func createTable(db *sql.DB) error {
	if _, err := db.Exec(queryCreateTasksTable); err != nil {
		return ErrFailedToCreateTable
	}

	return nil
}
