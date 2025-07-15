package account

import (
	"context"
	monimeapis2 "github.com/monimesl/monime-cli/cli-utils/monimeapis"
	"github.com/monimesl/monime-cli/pkg/monimeapis"
)

func init() {
	monimeapis.GetActiveAccountTokenFunc = getActiveToken

}

func getActiveToken(ctx context.Context) (string, error) {
	svc, err := NewService()
	if err != nil {
		return "", err
	}
	acc, ok, err := svc.GetActiveAccount(ctx)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", monimeapis2.ErrNotAuthenticated
	}
	return acc.Token, nil
}
