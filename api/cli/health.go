package handler

import (
	"api-authenticator-proxy/util/log"
	"github.com/urfave/cli/v2"
)

func healthCLI() *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "Check the health of the API Authenticator Proxy",
		Action: func(c *cli.Context) error {
			return health()
		},
	}
}

func health() error {
	log.CLI("Checking Http Server Status ...")
	log.CLI("Http Server is running ...")
	return nil
}
