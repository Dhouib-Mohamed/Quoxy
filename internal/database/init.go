package database

import (
	"api-authenticator-proxy/util/log"
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
	"time"
)

var db *sql.DB

func InitDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetConnMaxIdleTime(time.Minute * 5)

	// Read and execute the initialization script
	data, err := os.ReadFile("init.sql")
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
