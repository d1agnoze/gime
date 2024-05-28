package cmd

import (
	"gime/app"
	"os"
)

func InputHanlder() error {
	if isInputFromPipe() {
		return ReadFromPipe(os.Stdin, &media.Link)
	} else {
		file, e := getFile()
		if e != nil {
			return e
		}
		defer file.Close()
		return ReadFromFile(file, &media.Link)
	}
}

func RunGUI() {
	app := app.NewApp(&AssetMap.Asset, &media)
	app.Run()
}
