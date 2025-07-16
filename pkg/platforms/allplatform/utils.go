package allplatform

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd) error {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	go func() {
		for stdoutScanner.Scan() {
			fmt.Println(stdoutScanner.Text())
		}
	}()
	go func() {
		for stderrScanner.Scan() {
			_, _ = fmt.Fprintln(os.Stderr, stderrScanner.Text())
		}
	}()
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command exited with error: %w", err)
	}
	return nil
}
