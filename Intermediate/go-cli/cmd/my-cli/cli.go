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
	app.Name = "Network Testing"
	app.Usage = "Query your Websites to get more informataion"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "url",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks Up the IP addresses for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("url"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
