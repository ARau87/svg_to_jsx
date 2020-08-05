package main

import (
	"log"
	"os"

	"github.com/ARau87/svg_to_jsx/actions"
	"github.com/urfave/cli/v2"
)

func SetupApp() *cli.App {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "convert",
				Aliases: []string{"c"},
				Usage:   "Convert the svg files to jsx",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "dir",
						Aliases:  []string{"d"},
						Usage:    "Path to the directory containing the svg files",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "out",
						Aliases:  []string{"o"},
						Usage:    "Path to the output directory",
						Required: true,
					},
				},
				Action: actions.ConvertFiles,
			},
		},
	}

	return app
}

func main() {

	app := SetupApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
