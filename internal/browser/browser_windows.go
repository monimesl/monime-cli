package browser

import (
	"errors"
	"os/exec"
)

func open(ctx context.Context, url string) (Command, error) {
	cmd, err := runCmd(ctx, "rundll32", "url.dll,FileProtocolHandler", url)
	if e, ok := err.(*exec.Error); ok && e.Err == exec.ErrNotFound {
		return nil, errors.New("rundll32 url.dll,FileProtocolHandler not found")
	}
	return cmd, err
}
