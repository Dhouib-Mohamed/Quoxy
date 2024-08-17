package handler

import (
	"api-authenticator-proxy/util/log"
	"github.com/urfave/cli/v2"
	"os"
)

func CLI() {
	app := &cli.App{
		Name:  "quoxy",
		Usage: "API Authenticator Proxy CLI",

		Commands: []*cli.Command{
			healthCLI(),
			versionCLI(),
			subscriptionCLI(),
			tokenCLI(),
		},
	}
	log.Fatal(app.Run(os.Args))
}
