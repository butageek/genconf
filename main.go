package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gocolly/colly"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "genconf",
		Usage: "config file generator",
		Commands: []*cli.Command{
			{
				Name:  "git",
				Usage: "generate git config files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Value:   ".gitignore",
						Usage:   "file to be generated",
					},
					&cli.StringFlag{
						Name:    "search",
						Aliases: []string{"s"},
						Value:   "",
						Usage:   "keywords to be searched when generating .gitignore file",
					},
				},
				Action: func(c *cli.Context) error {
					git(c.String("file"), c.String("search"))
					return nil
				},
			},
			{
				Name:  "py",
				Usage: "generate python config files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Value:   "setup.py",
						Usage:   "file to be generated",
					},
				},
				Action: func(c *cli.Context) error {
					py(c.String("file"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func git(file, searchStr string) {
	// set log level
	log.SetLevel(log.InfoLevel)
	// set log formatter
	log.SetFormatter(&log.TextFormatter{
		// show full timestamp
		FullTimestamp: true,
	})

	switch file {
	case ".gitignore", "gitignore":
		url := "https://www.gitignore.io/api/" + searchStr

		c := colly.NewCollector()

		// register response callback
		c.OnResponse(func(r *colly.Response) {
			log.Info("generating .gitignore")
			err := ioutil.WriteFile(".gitignore", r.Body, 0644)
			if err != nil {
				panic(err)
			}
		})

		log.Info("visiting ", url)
		c.Visit(url)
	default:
		log.Error("file is not recognized, exiting")
		os.Exit(1)
	}
}

func py(file string) {
	// set log level
	log.SetLevel(log.InfoLevel)
	// set log formatter
	log.SetFormatter(&log.TextFormatter{
		// show full timestamp
		FullTimestamp: true,
	})

	switch file {
	case "setup.py", "setup":
		url := "https://raw.githubusercontent.com/butageek/genconf/master/config/setup.py"
		c := colly.NewCollector()

		// register response callback
		c.OnResponse(func(r *colly.Response) {
			log.Info("generating setup.py")
			err := ioutil.WriteFile("setup.py", r.Body, 0644)
			if err != nil {
				panic(err)
			}
		})

		log.Info("visiting ", url)
		c.Visit(url)
	default:
		log.Error("file is not recognized, exiting")
		os.Exit(1)
	}
}
