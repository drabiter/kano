package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "kano"
	app.Usage = "cli interface for HummingBird API"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		ListWatching()
		return nil
	}

	app.Run(os.Args)
}
