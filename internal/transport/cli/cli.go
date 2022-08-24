package cliApp

import (
	convertmp4tomp3 "cli/mp4ToMp3/pkg/convertMp4ToMp3"
	"os"

	"github.com/urfave/cli"
)

type Cli struct {
	app *cli.App
}

func NewCliApp() *Cli {
	return &Cli{
		app: cli.NewApp(),
	}
}

func (c *Cli) Run() {
	c.app.Name = "Mp4ToMp3Converter"
	c.app.Usage = "Convert mp4 to mp3"
	c.app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "inputdir, inpd",
			Value: ".",
			Usage: "Set input dirrectory",
		},
		cli.StringFlag{
			Name:  "outputdir, outd",
			Value: "./mp3",
			Usage: "Set output dirrectory",
		},
	}
	c.app.Action = func(c *cli.Context) error {

		dir := c.GlobalString("inputdir")
		files, _ := os.ReadDir(dir)
		convertmp4tomp3.Run(files, dir)
		return nil
	}
}
