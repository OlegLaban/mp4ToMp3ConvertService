package main

import (
	convertmp4tomp3 "cli/mp4ToMp3/pkg/convertMp4ToMp3"
	"fmt"
	"os"
)

func main() {

}

func cli() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("You should add dir")
		return
	}
	dir := args[1]
	files, _ := os.ReadDir(dir)
	convertmp4tomp3.Run(files, dir)
}
