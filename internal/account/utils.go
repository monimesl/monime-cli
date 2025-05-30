package account

import (
	"context"
	"github.com/monimesl/monime-cli/pkg/errors"
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
		return "", errors.ErrAccountNotAuthenticated
	}
	return acc.Token, nil
}
