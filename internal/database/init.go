package database

import (
	"api-authenticator-proxy/util/config"
	"api-authenticator-proxy/util/log"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/minus5/gofreetds"
	_ "github.com/sijms/go-ora"
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
	external := config.IsDatabaseExternal()
	driver := "sqlite"
	source := dbPath
	if external {
		driver = config.GetDatabaseDriver()
		host := config.GetDatabaseHost()
		port := config.GetDatabasePort()
		user := config.GetDatabaseUser()
		password := config.GetDatabasePassword()
		source = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbPath)
	}
	var err error
	db, err = sql.Open(driver, source)
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
	err = TestDB()
	if err != nil {
		return err
	}
	return nil
}
