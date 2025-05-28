package errors

import "errors"

var (
	ErrNoActivateSpace         = errors.New("no activate space")
	ErrAccountNotAuthenticated = errors.New("no authenticated")
)
