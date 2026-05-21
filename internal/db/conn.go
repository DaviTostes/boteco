package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() error {
	path, exists := os.LookupEnv("DB_PATH")
	if !exists {
		return fmt.Errorf("DB_PATH not set")
	}

	var err error
	DB, err = sql.Open("sqlite", path)

	_, _ = DB.Exec("PRAGMA journal_mode=WAL;")

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		date TEXT NOT NULL UNIQUE
	);
	`)

	return err
}
