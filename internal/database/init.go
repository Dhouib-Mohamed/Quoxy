package database

import (
	"api-authenticator-proxy/util/log"
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

var db *sql.DB

func TestDB() error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	res := db.Ping()
	return res
}

func InitDatabase(dbPath string, initPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(initPath)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(context.Background(), string(data))
	if err != nil {
		return err
	}

	log.Info("Database initialized successfully")
	return nil
}
