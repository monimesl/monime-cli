package browser

import "context"

func open(ctx context.Context, url string) (Command, error) {
	return runCmd(ctx, "open", url)
}
