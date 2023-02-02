package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var language string
	var count int

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "boom",
		Version: "v19.99.0",
		Usage:   "make an explosive entrance",
		// UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Value:       "english",
				Usage:       "language for the greeting",
				Required:    false,
				Destination: &language,
				EnvVars:     []string{"LANG2", "APP_LANG1", "LEGACY_COMPAT_LANG1"},
			},
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Use a randomized port",
				Value:       2,
				DefaultText: "random",
				Action: func(ctx *cli.Context, v int) error {
					if v >= 65536 {
						return fmt.Errorf("Flag port value %v out of range[0-65535]", v)
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "foo",
				Aliases: []string{"f"},
				Usage:   "foo greeting",
				Count:   &count,
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE2`",
			},
			&cli.BoolFlag{
				Name:  "ginger-crouton",
				Usage: "is it in the soup?",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "add",
				Category: "template1",
				Aliases:  []string{"a"},
				Usage:    "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:     "complete",
				Category: "template1",
				Aliases:  []string{"c"},
				Usage:    "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
			{
				Name:  "short",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("serve:", cCtx.Bool("serve"))
					fmt.Println("option:", cCtx.Bool("option"))
					fmt.Println("message:", cCtx.String("message"))
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "someone"
			fmt.Println("count", count)
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			fmt.Println("lang:", cCtx.String("lang"))
			fmt.Println("language:", language)
			fmt.Println("port:", cCtx.Int("port"))
			fmt.Println("config:", cCtx.String("config"))
			fmt.Println("foo", cCtx.Bool("foo"))
			// if !cCtx.Bool("ginger-crouton") {
			// 	return cli.Exit("Ginger croutons are not in the soup", 86)
			// }
			return nil
		},
	}

	// sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
