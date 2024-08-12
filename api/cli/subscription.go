package handler

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/error_handler"
	"api-authenticator-proxy/util/log"
	"github.com/urfave/cli/v2"
)

var subscription = database.Subscription{}

func subscriptionCLI() *cli.Command {
	return &cli.Command{
		Name:  "subscription",
		Usage: "Subscription operations",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get all subscriptions",
				Action: func(c *cli.Context) error {
					return subscriptionGet()
				},
			},
			{
				Name:  "get-by-id",
				Usage: "Get subscription by id",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					return subscriptionById(id)
				},
			},
			{
				Name:  "get-by-name",
				Usage: "Get subscription by name",
				Action: func(c *cli.Context) error {
					name := c.Args().First()
					if name == "" {
						log.CLI("Name is required")
						return nil
					}
					return subscriptionByName(name)
				},
			},
			{
				Name:  "create",
				Usage: "Create a new subscription",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Name of the subscription",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "frequency",
						Usage:    "Frequency of the subscription",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "rate-limit",
						Usage:    "Rate limit of the subscription",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					newSubscription := models.CreateSubscription{
						Name:      c.String("name"),
						Frequency: c.String("frequency"),
						RateLimit: c.Int("rate-limit"),
					}
					return subscriptionCreate(newSubscription)
				},
			},
			{
				Name:  "update",
				Usage: "Update a subscription",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "Name of the subscription",
					},
					&cli.StringFlag{
						Name:  "frequency",
						Usage: "Frequency of the subscription",
					},
					&cli.IntFlag{
						Name:  "rate-limit",
						Usage: "Rate limit of the subscription",
					},
				},
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					updateSubscription := models.UpdateSubscription{
						Name:      c.String("name"),
						Frequency: c.String("frequency"),
						RateLimit: c.Int("rate-limit"),
					}
					return subscriptionUpdate(id, updateSubscription)
				},
			},
			{
				Name:  "disable",
				Usage: "Disable a subscription",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						log.CLI("ID is required")
						return nil
					}
					return subscriptionDisable(id)
				},
			},
		},
	}
}

func subscriptionById(id string) error {
	subscription, err := subscription.GetById(id)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscription found: ", subscription)
	return nil
}

func subscriptionGet() error {
	subscriptions, err := subscription.GetAll()
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscriptions found: ", subscriptions)
	return nil

}

func subscriptionByName(name string) error {
	subscription, err := subscription.GetByName(name)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscription found: ", subscription)
	return nil
}

func subscriptionCreate(newSubscription models.CreateSubscription) error {
	id, err := subscription.Create(&newSubscription)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscription created with id: ", id)
	return nil
}

func subscriptionUpdate(id string, updateSubscription models.UpdateSubscription) error {
	err := subscription.Update(id, &updateSubscription)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscription updated")
	return nil
}

func subscriptionDisable(id string) error {
	err := subscription.Disable(id)
	if err != nil {
		error_handler.CLIHandler(err)
		return nil
	}
	log.CLI("Subscription disabled")
	return nil

}
