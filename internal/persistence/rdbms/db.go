package rdbms

import "database/sql"

func NewSqliteDB() (*sql.DB, error) {
	db, err := sql.Open("duckdb", "memory.db")
	if err != nil {
		return nil, err
	}
	err = createMemoryTable(db)
	if err != nil {
		return nil, err
	}
	err = createConversationTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createMemoryTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS memories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid UUID NOT NULL,
			conversation_id UUID NOT NULL,
			query TEXT NOT NULL,
			response TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func createConversationTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS conversations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid UUID NOT NULL,
			agent TEXT NOT NULL,
			user TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
