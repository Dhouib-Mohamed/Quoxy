package tests

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/error_handler"
	"api-authenticator-proxy/util/log"
	"fmt"
	"testing"
)

func init() {
	log.SetLogLevel(env.DEBUG)
}

func validateError(t *testing.T, res error_handler.StatusError, status int) {
	if res == nil {
		return
	}
	code, err := res.GetError()
	if code != 200 {
		if code != status {
			log.Error(fmt.Errorf("%s - Status : %d", err, code))
			t.Errorf("Expected Code is %d but we found %d : %s", status, code, err)
			return
		}
	}
	t.Logf("This is a success case - Status : %d", code)
}

func testDatabase(t *testing.T) {
	// Initialize the database
	err := database.InitDatabase("db.test.sqlite", "../../scripts/sql/")
	if err != nil {
		t.Fatal(err)
	}
	err = database.TestDB()
	if err != nil {
		t.Fatal(err)
	}
}
