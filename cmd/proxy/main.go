package main

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/proxy"
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/log"
)

func main() {
	log.Fatal(database.InitDatabase(env.GetDatabasePath(), env.GetDatabaseInitFile()))
	proxy.Proxy()
}
