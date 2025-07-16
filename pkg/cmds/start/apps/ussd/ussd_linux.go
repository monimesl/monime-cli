package ussd

import (
	"context"
	"fmt"
	"github.com/monimesl/monime-cli/pkg/platforms/allplatform"
	"github.com/monimesl/monime-cli/pkg/platforms/linux"
	"github.com/monimesl/monime-cli/pkg/text"
	"os"
	"path/filepath"
)

const (
	binaryName      = "monime-715"
	binaryAssetName = "monime-715-linux-amd64.tar.gz"
)

var (
	binaryInstalledDir = filepath.Join(os.Getenv("HOME"), ".local", "bin")
)

func launchApp(ctx context.Context) (*os.Process, error) {
	process, err, found := launchBinary(ctx)
	if found {
		return process, err
	}
	formattedAppName := text.Format(binaryName, text.FormatOptions{
		Color: "green",
		Bold:  true,
	})
	text.PrintStart("Installing '%s'\n", formattedAppName)
	err = linux.InstallReleaseBinary(ctx, binaryAssetName, binaryName, binaryInstalledDir)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ℹ️ Run this command again after following the above instruction.")
	os.Exit(0)
	return nil, nil
}

func launchBinary(ctx context.Context) (*os.Process, error, bool) {
	path := linux.IsBinaryInstalled(binaryName)
	if path != "" {
		err := allplatform.RunStringCommand(ctx, path)
		return nil, err, true
	}
	fmt.Printf("The binary '%s' is not installed.\n", binaryName)
	return nil, nil, false
}
