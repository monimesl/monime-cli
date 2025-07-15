package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/monime-lab/gwater"
	"github.com/monimesl/monime-cli/pkg/utils/monimeapis"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"os"
	"time"
	"ussd-simulator/backend/ussdgateway"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var (
	instanceId = gwater.UUID5("Monime *715#")
)

func main() {
	parseArguments()
	gateway := ussdgateway.New()
	err := wails.Run(&options.App{
		Title:       "Monime *715#",
		Width:       400,
		Height:      820,
		AlwaysOnTop: true,
		OnStartup:   gateway.Initialize,
		Bind: []interface{}{
			gateway,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		ErrorFormatter:   func(err error) any { return err.Error() },
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: instanceId,
		},
		Windows: &windows.Options{
			ResizeDebounceMS: 2,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			ProgramName:         "Monime *715#",
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Icon:    icon,
				Title:   "Monime *715#",
				Message: fmt.Sprintf("© %d Monime", time.Now().Year()),
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed start: %s", err)
	}
}

func parseArguments() {
	monimeapis.GetActiveAccountTokenFunc = func(ctx context.Context) (string, error) {
		tk := os.Getenv("MONIME_CLI_TOKEN")
		return tk, nil
	}
}
