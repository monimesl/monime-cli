package linux

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/monime-lab/gwater/ioutil"
	"github.com/monimesl/monime-cli/internal/platforms/allplatform"
	text2 "github.com/monimesl/monime-cli/internal/text"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func IsBinaryInstalled(binaryName string) string {
	if path, err := exec.LookPath(binaryName); err == nil {
		return path
	}
	return ""
}

func InstallReleaseBinary(ctx context.Context, assetName string, binaryName string, installDir string) error {
	tempDir, err := os.MkdirTemp("", "monime-cli-installer-")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()
	downloadURL, err := getLatestReleaseAssetURL(ctx, assetName)
	if err != nil {
		return fmt.Errorf("error getting download URL: %w", err)
	}
	archivePath := filepath.Join(tempDir, assetName)
	if err = downloadFile(ctx, archivePath, downloadURL); err != nil {
		return fmt.Errorf("error downloading archive: %w", err)
	}
	extractedBinaryPath, err := extractArchive(archivePath, tempDir, binaryName)
	if err != nil {
		return fmt.Errorf("error extracting archive: %w", err)
	}
	if err = installBinary(extractedBinaryPath, installDir, binaryName); err != nil {
		return fmt.Errorf("error installing binary: %w", err)
	}
	fmt.Println("\n--- Installation Complete! ---")
	fmt.Printf("The '%s' binary is now available at %s.\n", binaryName, filepath.Join(installDir, binaryName))
	fmt.Println("\nIMPORTANT: You may need to add this directory to your shell's PATH environment variable.")
	fmt.Println("Typically, you can do this by adding the following line to your shell's configuration file")
	fmt.Println("(.bashrc, .zshrc, .profile, etc.) and then restarting your terminal or sourcing the file:")
	fmt.Printf("  export PATH=\"$PATH:%s\"\n", installDir)
	fmt.Println("-----------------------------------------")
	return nil
}

func getLatestReleaseAssetURL(ctx context.Context, assetName string) (string, error) {
	type GithubRelease struct {
		Assets []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}
	apiURL := "https://api.github.com/repos/monimesl/monime-cli/releases/latest"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}
	//nolint:bodyclose // already handled by ioutil.CloseQuietly
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer ioutil.CloseQuietly(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch latest release: %w, status: %s", err, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("asset release fetch error: %w", err)
	}
	var release GithubRelease
	if err = json.Unmarshal(body, &release); err != nil {
		return "", err
	}
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			return asset.BrowserDownloadURL, nil
		}
	}
	return "", fmt.Errorf("asset '%s' not found in the latest release", assetName)
}

func downloadFile(ctx context.Context, filepath string, url string) error {
	text2.PrintStart("Downloading %s", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	//nolint:bodyclose // already handled by ioutil.CloseQuietly
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download asset file: %w", err)
	}
	defer ioutil.CloseQuietly(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status: %s", resp.Status)
	}
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filepath, err)
	}
	defer ioutil.CloseQuietly(out)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}
	return nil
}

func extractArchive(archivePath, extractDir, binaryName string) (string, error) {
	text2.PrintStart("Extracting the downloaded archive")
	tarPath, err := exec.LookPath("tar")
	if err != nil {
		return "", fmt.Errorf("'tar' command not found in PATH: %w", err)
	}
	if err := os.MkdirAll(extractDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create extraction directory %s: %w", extractDir, err)
	}
	cmd := exec.Command(tarPath, "-xzf", archivePath, "-C", extractDir)
	if err = allplatform.RunCommand(cmd); err != nil {
		return "", fmt.Errorf("failed to extract archive with 'tar' command: %w", err)
	}
	extractedBinaryPath := filepath.Join(extractDir, binaryName)
	if _, err := os.Stat(extractedBinaryPath); os.IsNotExist(err) {
		return "", fmt.Errorf("binary not found at %s", extractedBinaryPath)
	} else if err != nil {
		return "", fmt.Errorf("error checking extracted binary: %w", err)
	}
	return extractedBinaryPath, nil
}

func installBinary(sourcePath, installDir, binaryName string) error {
	text2.PrintStart("Installing extracted binary '%s' to %s...", binaryName, installDir)
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create installation directory %s: %w", installDir, err)
	}
	destPath := filepath.Join(installDir, binaryName)
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open extracted binary %s: %w", sourcePath, err)
	}
	defer ioutil.CloseQuietly(sourceFile)
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination binary file %s: %w", destPath, err)
	}
	defer ioutil.CloseQuietly(destFile)
	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy source binary: %w", err)
	}
	if err = os.Chmod(destPath, 0755); err != nil {
		return fmt.Errorf("failed to make the destination binary executable: %w", err)
	}
	return nil
}
