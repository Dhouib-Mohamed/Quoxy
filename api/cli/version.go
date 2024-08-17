package handler

import (
	"api-authenticator-proxy/util/log"
	"github.com/urfave/cli/v2"
	"os"
)

func versionCLI() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Get the version of the API Authenticator Proxy",
		Action: func(c *cli.Context) error {
			return version()
		},
	}
}

func version() error {
	version, _ := os.ReadFile("version.txt")
	log.CLI("API Authenticator Proxy version: ", string(version))
	return nil
}
