package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `mydocker is simple docker`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(ctx *cli.Context) error {
		// logs as json formatter
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
