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
	"strings"
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
	var source string
	if external {
		driver = config.GetDatabaseDriver()
		host := config.GetDatabaseHost()
		port := config.GetDatabasePort()
		user := config.GetDatabaseUser()
		password := config.GetDatabasePassword()
		database := config.GetDatabaseName()
		switch driver {
		case "sqlite":
			source = dbPath
		case "mysql":
			source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
		case "postgres":
			source = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
		case "mssql":
			source = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", host, user, password, port, database)
		default:
			log.Fatal(fmt.Errorf("database driver not supported"))
		}
	}
	var err error
	db, err = sql.Open(driver, source)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(initPath + "/init-" + driver + ".sql")
	if err != nil {
		return err
	}

	commands := strings.Split(string(data), ";")
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if len(command) > 0 {
			_, err = db.ExecContext(context.Background(), command)
			if err != nil {
				return err
			}
		}
	}

	log.Info("Database initialized successfully")
	err = TestDB()
	if err != nil {
		return err
	}
	return nil
}
