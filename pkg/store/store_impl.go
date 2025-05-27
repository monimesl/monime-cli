package store

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
	"os"
	"path/filepath"
)

const (
	serviceName = "monime"
)

func newDefaultStore() (Store, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", serviceName)
	configFile := filepath.Join(configDir, "config.yaml")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err = os.WriteFile(configFile, []byte{}, 0600); err != nil {
			return nil, err
		}
	}
	v := viper.New()
	v.AddConfigPath(configDir)
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	keyringSvcName := fmt.Sprintf("io.monime.apps/cli/%s", serviceName)
	return &defaultStore{viper: v, keyringServiceName: keyringSvcName}, nil
}

type defaultStore struct {
	keyringServiceName string
	viperStoragePath   string
	viper              *viper.Viper
}

func (s *defaultStore) SetValue(key string, value any) error {
	return s.writeValue(key, value)
}

func (s *defaultStore) GetString(key string) string {
	return s.viper.GetString(key)
}

func (s *defaultStore) GetBool(key string) bool {
	return s.viper.GetBool(key)
}

func (s *defaultStore) GetInt(key string) int64 {
	return s.viper.GetInt64(key)
}

func (s *defaultStore) GetFloat64(key string) float64 {
	return s.viper.GetFloat64(key)
}

func (s *defaultStore) GetSecret(key string) (string, error) {
	return keyring.Get(s.keyringServiceName, key)
}

func (s *defaultStore) SetSecret(key, secret string) error {
	return keyring.Set(s.keyringServiceName, key, secret)
}

func (s *defaultStore) SetConfig(key string, value any) error {
	mp, err := s.convertToMap(value)
	if err != nil {
		return err
	}
	return s.SetValue(key, mp)
}

func (s *defaultStore) GetConfig(key string, out any) error {
	v := s.viper.Sub(key)
	if v == nil {
		return ErrKeyNotFound
	}
	if err := v.Unmarshal(out); err != nil {
		return err
	}
	return nil
}

func (s *defaultStore) writeValue(key string, value any) error {
	s.viper.Set(key, value)
	return s.viper.WriteConfig()
}

func (s *defaultStore) convertToMap(value any) (map[string]any, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	m := map[string]any{}
	if err = json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}
	return m, nil
}
