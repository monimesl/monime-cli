package store

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Store interface {
	DeleteSecret(key string) error
	GetSecret(key string) (string, error)
	SetSecret(key, secret string) error
	SetConfig(key string, value any) error
	GetConfig(key string, out any) error
	SetValue(key string, value any) error
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int64
	GetFloat64(key string) float64
}
