package browser

import (
	"os"
	"os/exec"
)

func OpenURL(url string) (*exec.Cmd, error) {
	return open(url)
}

func runCmd(cmd string, args ...string) (Command, error) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return nil, err
	}
	return c, nil
}

type Command *exec.Cmd
