package tests

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/util/error_handler"
	"fmt"
	"testing"
)

func validateError(t *testing.T, res error_handler.StatusError, status int) {
	if res == nil {
		return
	}
	code, err := res.GetError()
	if code != 200 {
		if code != status {
			t.Fatal(fmt.Sprintf("Expected Code is %d but we found %d : %s", status, code, err))
			return
		}
		t.Log(fmt.Sprintf("%s - Status : %d", err, code))
	}
	t.Log(fmt.Sprintf("This is a success case - Status : %d", code))
}

func testDatabase(t *testing.T) {
	// Initialize the database
	err := database.InitDatabase("db.test.sqlite", "../../init.sql")
	if err != nil {
		t.Error(err)
	}
	err = database.TestDB()
	if err != nil {
		t.Error(err)
	}
}
