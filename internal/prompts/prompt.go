package prompts

import (
	"context"
	"os"

	"golang.org/x/term"
)

func EnterKey(ctx context.Context) (bool, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return false, err
	}
	defer func(fd int, oldState *term.State) {
		_ = term.Restore(fd, oldState)
	}(fd, oldState)
	resultChan := make(chan byte, 1)
	errChan := make(chan error, 1)
	defer func() {
		close(errChan)
		close(resultChan)
	}()
	go func() {
		buf := make([]byte, 1)
		if _, err := os.Stdin.Read(buf); err != nil {
			errChan <- err
			return
		}
		resultChan <- buf[0]
	}()
	select {
	case <-ctx.Done():
		return false, nil
	case err = <-errChan:
		return false, err
	case buf := <-resultChan:
		return buf == '\r' || buf == '\n', nil
	}
}
