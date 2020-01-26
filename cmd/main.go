package main

import (
	"log"
	"os"

	"github.com/PumpkinSeed/kcs"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "kcs",
		Usage: "Searchable kubectl cheatsheet CLI tool",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Set verbose mode",
			},
			&cli.StringFlag{
				Name: "search",
				Aliases: []string{"s"},
				Usage: "Set search query",
			},
			&cli.StringFlag{
				Name: "category",
				DefaultText: "",
				Usage: "Set a certain category",
			},
			&cli.StringFlag{
				Name: "command",
				DefaultText: "",
				Usage: "Set a certain command",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("verbose") {
				kcs.SetVerbose(true)
			}
			if q := c.String("search"); q != "" {
				kcs.PrintSearchResult(kcs.Search(q))
				return nil
			}
			kcs.Data.Print(c.String("category"), c.String("command"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
