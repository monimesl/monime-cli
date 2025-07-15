package account

import (
	"context"
	"fmt"
	"github.com/monimesl/monime-cli/cli-utils/monimeapis"
	"net/http"
)

func init() {
	monimeapis.GetActiveAccountTokenFunc = getActiveToken
}

func CheckActiveToken(ctx context.Context) (string, error) {
	token, err := getActiveToken(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println("⌛️ Checking validity of authentication context")
	_, err = monimeapis.ApiRequest[any, any](ctx, nil, http.MethodPost, "/tokens/check", nil)
	if err != nil {
		return "", err
	}
	return token, nil
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
		return "", monimeapis.ErrNotAuthenticated
	}
	return acc.Token, nil
}
