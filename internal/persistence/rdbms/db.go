package rdbms

import (
	"database/sql"

	_ "github.com/duckdb/duckdb-go/v2"
)

func NewDuckDB() (*sql.DB, error) {
	db, err := sql.Open("duckdb", "memory.db")
	if err != nil {
		return nil, err
	}
	err = createSequences(db)
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
	err = createMemoryMetaTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createSequences(db *sql.DB) error {
	queries := []string{
		"CREATE SEQUENCE IF NOT EXISTS memories_id_seq START 1",
		"CREATE SEQUENCE IF NOT EXISTS conversations_id_seq START 1",
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func createMemoryTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS memories (
			id INTEGER PRIMARY KEY DEFAULT nextval('memories_id_seq'),
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

func createMemoryMetaTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS memories_meta (
			id INTEGER PRIMARY KEY DEFAULT nextval('memories_id_seq'),
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
			id INTEGER PRIMARY KEY DEFAULT nextval('conversations_id_seq'),
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
