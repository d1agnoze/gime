package main

import (
	"fmt"
	i "gime/internal"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) HelloWorld() string {
	return fmt.Sprint("Yeet")
}

func (a *App) GetMedia() i.Media {
	return i.Media{
		Mime: a.media.Mime,
		Link: a.media.Link,
	}
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
}

