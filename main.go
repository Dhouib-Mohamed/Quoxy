package main

import (
	"api-authenticator-proxy/src/database"
	"api-authenticator-proxy/src/proxy"
	"api-authenticator-proxy/src/router"
)

func main() {
	go func() {
		err := database.Database("db.sqlite")
		if err != nil {
			panic(err)
		}
	}()
	go router.Router()
	proxy.Proxy()
}
