package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go cli app"
	app.Usage = "Search public IP or server name by host using cli"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search addresses ips on web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) {
				hostParam := c.String("host")

				ips, err := net.LookupIP(hostParam)
				if err != nil {
					log.Fatal("Get ips error: ", err)
				}

				for _, ip := range ips {
					fmt.Println(ip)
				}
			},
		},
	}

	app.Run(os.Args)
}
