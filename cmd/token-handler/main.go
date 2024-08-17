package main

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/frequency_cron"
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/log"
)

func main() {
	log.Fatal(database.InitDatabase(env.GetDatabasePath(), env.GetDatabaseInitFile()))
	frequency_cron.Init()
}
