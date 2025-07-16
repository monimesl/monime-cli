package darwin

import (
	"context"
	"fmt"
	"github.com/monimesl/monime-cli/internal/platforms/allplatform"
	"os"
	"os/exec"
	"path/filepath"
)

func IsAppInstalled(appName string) (bool, error) {
	appPath := filepath.Join("/Applications", appName+".app")
	info, err := os.Stat(appPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("error checking path %s: %w", appPath, err)
	}
	if info.IsDir() {
		return true, nil
	}
	return false, fmt.Errorf("path %s exists but is not a .app directory", appPath)
}

func OpenApp(ctx context.Context, appName, binaryName string) (*os.Process, error) {
	appPath := filepath.Join("/Applications", appName+".app")
	appExecutablePath := filepath.Join(appPath, "Contents", "MacOS", binaryName)
	cmd := exec.CommandContext(ctx, appExecutablePath)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to start app %s: %w", appName, err)
	}
	return cmd.Process, nil
}

func InstallBrewCask(ctx context.Context, caskName, tapRepo string) error {
	brewPath, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("--- Homebrew Not Found ---")
		fmt.Println("Homebrew is required but not found on your system.")
		fmt.Println("Please install Homebrew first by following instructions at https://brew.sh/")
		fmt.Println("--------------------------")
		return fmt.Errorf("homebrew not found — please install it from https://brew.sh/: %w", err)
	}
	if tapRepo != "" {
		cmd := exec.Command(brewPath, "tap", tapRepo)
		if err = allplatform.RunCommand(cmd); err != nil {
			return fmt.Errorf("failed to tap repository %s: %w", tapRepo, err)
		}
	}
	cmd := exec.CommandContext(ctx, brewPath, "install", "--cask", caskName)
	if err = allplatform.RunCommand(cmd); err != nil {
		fmt.Printf("Failed to install cask %s: %v\n", caskName, err)
		return fmt.Errorf("failed to install Homebrew cask: %w", err)
	}
	return nil
}
