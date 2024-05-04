package main

import (
	"api-authenticator-proxy/src/database"
	"api-authenticator-proxy/src/frequency_cron"
	"api-authenticator-proxy/src/proxy"
	"api-authenticator-proxy/src/router"
)

func main() {
	go func() {
		err := database.InitDatabase("db.sqlite")
		if err != nil {
			panic(err)
		}
	}()
	go frequency_cron.Init()
	go router.Router()
	proxy.Proxy()
}
