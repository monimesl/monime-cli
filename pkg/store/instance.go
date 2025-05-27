package store

import (
	"log"
	"sync"
)

type Provider func() (Store, error)

var (
	DefaultProvider Provider = newDefaultStore
	instance        Store
	once            sync.Once
)

func Get() Store {
	once.Do(func() {
		if instance == nil {
			st, err := DefaultProvider()
			if err != nil {
				log.Fatal(err)
			}
			instance = st
		}
	})
	return instance
}
