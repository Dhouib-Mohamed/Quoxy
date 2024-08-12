package handler

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/error_handler"
	"api-authenticator-proxy/util/log"
	"github.com/urfave/cli/v2"
)

var token = database.Token{}

func tokenCLI() *cli.Command {
	return &cli.Command{
		Name:  "token",
		Usage: "Token operations",
		Subcommands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get all tokens",
				Action: func(c *cli.Context) error {
					return tokenGet()
				},
			},
			{
				Name:  "get-by-id",
				Usage: "Get token by id",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					return tokenById(id)
				},
			},
			{
				Name:  "create",
				Usage: "Create a new token",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "passphrase",
						Usage: "Passphrase of the token",
					},
					&cli.StringFlag{
						Name:     "subscription",
						Usage:    "Subscription of the token",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					newToken := models.CreateToken{
						Passphrase:   c.String("passphrase"),
						Subscription: c.String("subscription"),
					}
					return tokenCreate(newToken)
				},
			},
			{
				Name:  "update",
				Usage: "Update a token",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "subscription",
						Usage:    "Subscription of the token",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					updateToken := models.UpdateToken{
						Subscription: c.String("subscription"),
					}
					return tokenUpdate(id, updateToken)
				},
			},
			{
				Name:  "disable",
				Usage: "Disable a token",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					return tokenDisable(id)
				},
			},
		},
	}
}

func tokenGet() error {
	tokens, err := token.GetAll()
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	if len(tokens) == 0 {
		log.CLI("No tokens found")
		return nil
	}
	log.CLI("Tokens:")
	for _, token := range tokens {
		log.CLI(token)
	}
	return nil

}

func tokenById(id string) error {
	token, err := token.GetById(id)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI(token)
	return nil

}

func tokenCreate(newToken models.CreateToken) error {
	id, err := token.Create(&newToken)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Token created with id: ", id)
	return nil

}

func tokenUpdate(id string, updateToken models.UpdateToken) error {
	err := token.Update(id, &updateToken)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Token updated")
	return nil

}

func tokenDisable(id string) error {
	err := token.Disable(id)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Token deleted")
	return nil
}
