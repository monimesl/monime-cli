package ussd

import (
	"context"
	"github.com/monimesl/monime-cli/pkg/platforms/darwin"
	"github.com/monimesl/monime-cli/pkg/text"
	"os"
)

const (
	appName      = "Monime 715"
	brewCaskName = "monime-715"
	brewTapRepo  = "monimesl/monime-cli"
)

func launchApp(ctx context.Context) (*os.Process, error) {
	process, err, appExists := launchBrewCaskApp(ctx)
	if appExists {
		return process, err
	}
	formattedAppName := text.Format(appName, text.FormatOptions{
		Color: "green",
		Bold:  true,
	})
	text.PrintStart("Installing %s using brew\n", formattedAppName)
	if err = darwin.InstallBrewCask(ctx, brewCaskName, brewTapRepo); err != nil {
		return nil, err
	}
	text.PrintStart("Launching the installed %s\n", formattedAppName)
	process, err, _ = launchBrewCaskApp(ctx)
	return process, err
}

func launchBrewCaskApp(ctx context.Context) (*os.Process, error, bool) {
	if yes, err := darwin.IsAppInstalled(appName); err != nil {
		return nil, err, true
	} else if yes {
		process, err := darwin.OpenApp(ctx, appName, "monime-715")
		return process, err, true
	}
	return nil, nil, false
}
