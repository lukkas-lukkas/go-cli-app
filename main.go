package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go cli app"
	app.Usage = "Search public IP or server name by host using cli"

	app.Run(os.Args)
}
