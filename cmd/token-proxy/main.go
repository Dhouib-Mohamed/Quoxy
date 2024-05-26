package main

import (
	"api-authenticator-proxy/api/handler"
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/frequency_cron"
	"api-authenticator-proxy/internal/proxy"
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/log"
)

func main() {
	log.Fatal(database.InitDatabase(env.GetDatabasePath(), env.GetDatabaseInitFile()))
	go frequency_cron.Init()
	go handler.Router()
	proxy.Proxy()
}
