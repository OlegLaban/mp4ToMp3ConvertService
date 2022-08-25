package main

import (
	cliApp "cli/mp4ToMp3/internal/transport/cli"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		cliApp.NewCliApp().Run(os.Args)
	}

}
