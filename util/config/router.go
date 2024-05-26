package config

import (
	"api-authenticator-proxy/util/log"
)

type RouterEnv struct {
	Disabled bool   `yaml:"disabled"`
	Port     string `yaml:"port"`
}

var router RouterEnv

func init() {
	router = RouterEnv{}
	getConfigVar(&router, "router")
}

func GetRouterPort() string {
	return getValidPort(router.Port)
}

func GetIsRouterEnabled() bool {
	if router.Disabled == true {
		log.Warning("Router is disabled. Please enable the router in config.yaml")
	}
	return !router.Disabled
}
