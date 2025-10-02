package browser

import (
	"context"
	"os"
	"os/exec"
)

func OpenURL(ctx context.Context, url string) (*exec.Cmd, error) {
	return open(ctx, url)
}

func runCmd(ctx context.Context, cmd string, args ...string) (Command, error) {
	c := exec.CommandContext(ctx, cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return nil, err
	}
	return c, nil
}

type Command *exec.Cmd
