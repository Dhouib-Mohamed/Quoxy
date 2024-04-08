package database

import (
	"database/sql"
	"fmt"
)

func checkResponse(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("Unexpected error, try checking the provided id")
	}
	return nil
}

func GetLastInsertedId(table string) (string, error) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY ROWID DESC LIMIT 1", table)
	row := db.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
