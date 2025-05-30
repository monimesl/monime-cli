package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "Monime Webhook",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
		},
		Mac: &mac.Options{
			Appearance: mac.NSAppearanceNameVibrantLight,
			TitleBar: &mac.TitleBar{
				FullSizeContent: true,
			},
			About: &mac.AboutInfo{
				Title:   "Monime Webhook",
				Message: "A webhook monitor for testing and debugging Monime webhooks api.",
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed start: %s", err)
	}
}
