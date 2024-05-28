package main

import (
	"context"
	"fmt"
	i "gime/internal"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const TITLE = "gime"
const DEFAULT_W = 520
const DEFAULT_H = 300
const MIN_W = 360
const MIN_H = 200

// App struct
type App struct {
	ctx   context.Context
	media i.Media
}

func NewApp(m *i.Media) *App {
	return &App{media: *m}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	screen, err := getScreen(ctx)
	if err != nil {
		return
	}
	size_X, size_Y := runtime.WindowGetSize(a.ctx)
	appSize := i.Coordinate{X: size_X, Y: size_Y}

	appPostion := i.Position{Name: position}
	appCoordinate := appPostion.GetCoordinate(appSize, *screen)
	if appCoordinate == nil {
		return
	}
	runtime.WindowSetPosition(a.ctx, appCoordinate.X, appCoordinate.Y)
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
			Assets: assets,
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

func getScreen(ctx context.Context) (*i.Coordinate, error) {
	screens, err := runtime.ScreenGetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, screen := range screens {
		if screen.IsPrimary {
			return &i.Coordinate{X: screen.PhysicalSize.Width, Y: screen.PhysicalSize.Height}, nil
		}
	}
	return nil, fmt.Errorf("no primary screen found")
}
