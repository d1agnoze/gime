package main

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// return type for the front-end
type Media struct {
	Mime string
	Link string
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMedia() Media {
	return Media{
		Mime: a.mime,
		Link: a.link,
	}
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
}
