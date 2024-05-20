package database

import (
	error_handler2 "api-authenticator-proxy/util/error_handler"
	dbError "api-authenticator-proxy/util/error_handler/db"
	"database/sql"
	"fmt"
)

func checkWriteResponse(result sql.Result, err error, table string) error_handler2.StatusError {
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			return dbError.ElementNotFoundError(table)
		default:
			return error_handler2.UnexpectedError(fmt.Sprintf("Unexpected error when writing to the %s table ", table))
		}
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return error_handler2.UnexpectedError(fmt.Sprintf("Unexpected error when writing to the %s table ", table))
	}
	if rows == 0 {
		return dbError.ElementNotFoundError(table)
	}
	return nil
}

func checkReadResponse(err error, table string) error_handler2.StatusError {
	if err == nil {
		return nil
	}
	switch err.Error() {
	case sql.ErrNoRows.Error():
		return dbError.ElementNotFoundError(table)
	default:
		return error_handler2.UnexpectedError(fmt.Sprintf("Unexpected error when reading from the %s table ", table))
	}
}

func GetLastInsertedId(table string) (string, error_handler2.StatusError) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY ROWID DESC LIMIT 1", table)
	row := db.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		return "", error_handler2.UnexpectedError(fmt.Sprintf("Unexpected error when reading from the %s table : %s", table, err.Error()))
	}
	return id, nil
}
