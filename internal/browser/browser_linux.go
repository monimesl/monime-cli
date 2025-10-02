package browser

import (
	"os/exec"
	"strings"
)

func open(ctx context.Context, url string) (Command, error) {
	programs := []string{"xdg-open", "x-www-browser", "www-browser"}
	for _, provider := range programs {
		if _, err := exec.LookPath(provider); err == nil {
			return runCmd(ctx, provider, url)
		}
	}
	return nil, &exec.Error{Name: strings.Join(programs, ","), Err: exec.ErrNotFound}
}
