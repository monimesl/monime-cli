package darwin

import (
	"context"
	"errors"
	"fmt"
	"github.com/monimesl/monime-cli/pkg/platforms/allplatform"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func IsAppInstalled(appName string) (bool, error) {
	if runtime.GOOS != "darwin" {
		return false, fmt.Errorf("platform must be darwin, got %s", runtime.GOOS)
	}
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
	if runtime.GOOS != "darwin" {
		return nil, fmt.Errorf("platform must be darwin, got %s", runtime.GOOS)
	}
	appPath := filepath.Join("/Applications", appName+".app")
	appExecutablePath := filepath.Join(appPath, "Contents", "MacOS", binaryName)
	cmd := exec.CommandContext(ctx, appExecutablePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, nil
		}
		return nil, fmt.Errorf("error opening app %s: %w\n%s", appName, err, output)
	}
	return cmd.Process, nil
}

func InstallBrewCask(ctx context.Context, caskName, tapRepo string) error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("brew cask installation is only supported on macOS")
	}
	brewPath, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("--- Homebrew Not Found ---")
		fmt.Println("Homebrew is required but not found on your system.")
		fmt.Println("Please install Homebrew first by following instructions at https://brew.sh/")
		fmt.Println("--------------------------")
		return fmt.Errorf("brew binary not found")
	}
	if tapRepo != "" {
		cmd := exec.Command(brewPath, "tap", tapRepo)
		if err := allplatform.RunCommand(cmd); err != nil {
			return fmt.Errorf("Failed to tap repository %s. It might already be tapped or an error occurred: %v\n", tapRepo, err)
		}
	}
	cmd := exec.CommandContext(ctx, brewPath, "install", "--cask", caskName, "-q")
	if err := allplatform.RunCommand(cmd); err != nil {
		fmt.Printf("Failed to install cask %s: %v\n", caskName, err)
		return fmt.Errorf("failed to install Homebrew cask: %w", err)
	}
	return nil
}
