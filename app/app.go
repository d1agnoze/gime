package app

import (
	"context"
	"embed"
	i "gime/internal"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

const TITLE = "gime"
const DEFAULT_W = 520
const DEFAULT_H = 300
const MIN_W = 360
const MIN_H = 200

// App struct
type App struct {
	ctx    context.Context
	assets embed.FS
	media  i.Media
}

func NewApp(a *embed.FS, m *i.Media) *App {
	return &App{
		assets: *a,
		media:  *m,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
}

func (app *App) Run() {
	err := wails.Run(&options.App{
		Title:     TITLE,
		Width:     DEFAULT_W,
		Height:    DEFAULT_H,
		MinWidth:  MIN_W,
		MinHeight: MIN_H,
		Frameless: true,
		Bind: []interface{}{
			app,
			&i.Media{},
		},
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		Windows: &windows.Options{
			Theme:               windows.Dark,
			BackdropType:        windows.Acrylic,
			WindowIsTranslucent: true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         TITLE,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
