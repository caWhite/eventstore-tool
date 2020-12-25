package main

import (
	"github.com/cawhite/eventstore-tool/backup/internal/ec2"
	"github.com/urfave/cli/v2"
	"os"
)

var app = &cli.App{
	Name: "backup",
	Description: "EventStoreDB backup utility",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "environment",
			Aliases: []string{"e"},
			Usage:   "Filter instances to `ENVIRONMENT`",
		},
	},
	Action: func(context *cli.Context) error {
		ec2.GetInstances(context.String("environment"))
		return nil
	},
}

func main()  {
	app.Run(os.Args)
}