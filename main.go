package main

import (
	"embed"
	"gime/cmd"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cmd.AssetMap = cmd.Assets{
		Asset: assets,
	}
	cmd.Excute()
}
