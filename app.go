package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "kano"
	app.Usage = "cli interface for HummingBird API"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("hola " + c.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
