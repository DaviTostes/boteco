package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() error {

	path := "/home/toast/.local/share/boteco/database.sqlite"

	var err error
	DB, err = sql.Open("sqlite", path)

	_, _ = DB.Exec("PRAGMA journal_mode=WAL;")
	_, _ = DB.Exec("PRAGMA busy_timeout = 5000;")

	DB.SetMaxOpenConns(1)

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		date TEXT NOT NULL UNIQUE
	);

	CREATE TABLE IF NOT EXISTS memories(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS chats(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS chat_messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		chat_id INTEGER NOT NULL,
		FOREIGN KEY(chat_id) REFERENCES chat(id)
	);
	`)

	return err
}
