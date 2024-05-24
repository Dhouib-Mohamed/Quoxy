package database

import (
	"api-authenticator-proxy/util/error_handler"
	dbError "api-authenticator-proxy/util/error_handler/db"
	"api-authenticator-proxy/util/log"
	"database/sql"
	"fmt"
	"strings"
)

func checkWriteResponse(result sql.Result, err error, table string) error_handler.StatusError {
	if err != nil {
		log.Error(fmt.Errorf("Error when writing to the %s table : %s", table, err.Error()))
		switch err.Error() {
		case sql.ErrNoRows.Error():
			log.Debug(fmt.Sprintf("No Element following this criteria was found in the %s table", table))
			return dbError.ElementNotFoundError(table)
		default:
			if strings.Contains(err.Error(), "constraint failed:") {
				log.Debug(err.Error())
				if strings.Contains(err.Error(), "UNIQUE constraint failed: ") {
					val1 := strings.Split(err.Error(), "UNIQUE constraint failed: ")[1]
					val2 := strings.Split(val1, " ")[0]
					field := strings.Split(val2, ".")[1]
					return dbError.FieldConstraintError(table, field, "should be UNIQUE")
				}
				return dbError.FieldConstraintError(table, "", "Unknown")
			}
			return error_handler.UnexpectedError(fmt.Sprintf("Unexpected error when writing to the %s table ", table))
		}
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Error(fmt.Errorf("Error when writing to the %s table : %s", table, err.Error()))
		return error_handler.UnexpectedError(fmt.Sprintf("Unexpected error when writing to the %s table ", table))
	}
	if rows == 0 {
		return dbError.ElementNotFoundError(table)
	}
	log.Debug(fmt.Sprintf("Successfully changed %d item in the %s table", rows, table))
	return nil
}

func checkReadResponse(err error, table string) error_handler.StatusError {
	if err != nil {
		log.Error(fmt.Errorf("Error when reading from the %s table : %s", table, err.Error()))
		switch err.Error() {
		case sql.ErrNoRows.Error():
			return dbError.ElementNotFoundError(table)
		default:
			return error_handler.UnexpectedError(fmt.Sprintf("Unexpected error when reading from the %s table ", table))
		}
	}
	log.Debug(fmt.Sprintf("Successfully read 1 item from the %s table", table))
	return nil
}

func GetLastInsertedId(table string) (string, error_handler.StatusError) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY ROWID DESC LIMIT 1", table)
	row := db.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		return "", error_handler.UnexpectedError(fmt.Sprintf("Unexpected error when reading from the %s table : %s", table, err.Error()))
	}
	return id, nil
}
