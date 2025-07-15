package monimeapis

import (
	"context"
	"errors"
)

type getActiveAccountTokenFunc func(ctx context.Context) (string, error)

var (
	GetActiveAccountTokenFunc getActiveAccountTokenFunc
)

func getActiveAccountToken(ctx context.Context) (string, error) {
	if GetActiveAccountTokenFunc != nil {
		return GetActiveAccountTokenFunc(ctx)
	}
	return "", errors.New("not implemented")
}
