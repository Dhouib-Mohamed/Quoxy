package config

import (
	"api-authenticator-proxy/util/log"
	"api-authenticator-proxy/util/network"
)

type ProxyEnv struct {
	Disabled bool   `yaml:"disabled"`
	Port     string `yaml:"port"`
	Target   string `yaml:"target"`
}

var proxy ProxyEnv

func init() {
	proxy = ProxyEnv{}
	log.Fatal(getConfigVar(&proxy, "proxy"))
}

func GetProxyPort() string {
	return getValidPort(proxy.Port)
}

func GetProxyTarget() string {
	defaultTarget := "https://google.com"
	if proxy.Target == "" {
		log.Warning("Proxy target not set. Please set the target in config.yaml")
		return defaultTarget
	}
	if !network.IsURLValid(proxy.Target) {
		log.Warning("Invalid proxy target: ", proxy.Target)
		return defaultTarget
	}
	return proxy.Target
}

func GetIsProxyEnabled() bool {
	if proxy.Disabled {
		log.Warning("Proxy is disabled. Please enable the router in config.yaml")
	}
	return !proxy.Disabled
}
