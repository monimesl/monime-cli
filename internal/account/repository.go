package account

import (
	"errors"
	"github.com/monimesl/monime-cli/pkg/store"
)

type Repository interface {
	ListAccounts() (List, error)
	SaveAccounts(list List) error
	AddAccount(account Account) error
	GetAccountById(id string) (Account, bool, error)
}

const (
	accountsField     = "accounts"
	accountIdField    = "account_id"
	accountTokenField = "account_token"
)

type defaultRepository struct{}

func (r defaultRepository) ListAccounts() (List, error) {
	list := List{}
	err := store.Get().GetConfig(accountsField, &list)
	if err != nil && !errors.Is(err, store.ErrKeyNotFound) {
		return List{}, err
	}
	return list, nil
}

func (r defaultRepository) SaveAccounts(list List) error {
	return store.Get().SetConfig(accountsField, list)
}

func (r defaultRepository) AddAccount(account Account) error {
	list, err := r.ListAccounts()
	if err != nil {
		return err
	}
	list.Add(account)
	if err = r.SaveAccounts(list); err != nil {
		return err
	}
	if err = store.Get().SetSecret(accountIdField, account.Id); err != nil {
		return err
	}
	if err = store.Get().SetSecret(accountTokenField, account.Token); err != nil {
		return err
	}
	return nil
}

func (r defaultRepository) GetAccountById(id string) (Account, bool, error) {
	list, err := r.ListAccounts()
	if err != nil {
		return Account{}, false, err
	}
	if acc, ok := list.GetById(id); ok {
		if acc.Id, err = store.Get().GetSecret(accountIdField); err != nil {
			return Account{}, false, err
		}
		if acc.Token, err = store.Get().GetSecret(accountTokenField); err != nil {
			return Account{}, false, err
		}
		return acc, true, nil
	}
	return Account{}, false, nil
}
