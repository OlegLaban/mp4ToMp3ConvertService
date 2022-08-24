package convertmp4tomp3

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const (
	VideoFormat = "mp4"
	AudioFormat = "mp3"
)

func Run(files []fs.DirEntry, dir string) {
	os.Mkdir(getFolderForAudioFormat(dir), 0755)
	c := make(chan string)
	j := tryConvert(files, dir, c)
	printSuccessMessagies(j, c)
}

func tryConvert(files []fs.DirEntry, dir string, c chan string) int {
	j := 0
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), "."+VideoFormat) {
			continue
		}
		go convert(dir, file, c)
		j++
	}

	return j
}

func printSuccessMessagies(j int, c chan string) {
	for i := 0; i < j; i++ {
		name := <-c
		fmt.Printf("File %s is converted\n", name)
	}
}

func getMp3Path(file fs.DirEntry, dir string) string {
	name := file.Name()
	return getFolderForAudioFormat(dir) + "/" + strings.ReplaceAll(name, "."+VideoFormat, "."+AudioFormat)
}

func getFolderForAudioFormat(dir string) string {
	return fmt.Sprintf("%s/%s", dir, AudioFormat)
}

func convert(dir string, file fs.DirEntry, c chan string) {
	ffmpeg.Input(dir+"/"+file.Name()).Output(getMp3Path(file, dir), ffmpeg.KwArgs{"map": "0:a:0", "b:a": "96k"}).OverWriteOutput().ErrorToStdOut().Run()
	c <- file.Name()
}
