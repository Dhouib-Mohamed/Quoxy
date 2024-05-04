package database

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
)

var db *sql.DB

func InitDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	data, err := os.ReadFile("init.sql")

	if err != nil {
		return err
	}

	_, err = db.ExecContext(context.Background(), string(data), nil)

	if err != nil {
		return err
	}
	return nil
}
