package main

import (
	"api-authenticator-proxy/src/database"
	"api-authenticator-proxy/src/frequency_cron"
	"api-authenticator-proxy/src/proxy"
	"api-authenticator-proxy/src/router"
	"api-authenticator-proxy/src/utils/log"
)

func main() {
	log.Fatal(database.InitDatabase("db.sqlite"))
	go frequency_cron.Init()
	go router.Router()
	proxy.Proxy()
}
