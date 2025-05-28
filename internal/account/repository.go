package account

import (
	"context"
	"errors"
	"github.com/monimesl/monime-cli/pkg/store"
)

type Repository interface {
	ListAccounts(context.Context) (List, error)
	SaveAccounts(ctx context.Context, list List) error
	AddAccount(ctx context.Context, account Account) error
	GetAccountById(ctx context.Context, id string) (Account, bool, error)
}

const (
	accountsField     = "accounts"
	accountIdField    = "account_id"
	accountTokenField = "account_token"
)

var (
	_ Repository = &defaultRepository{}
)

type defaultRepository struct{}

func (r defaultRepository) ListAccounts(context.Context) (List, error) {
	list := List{}
	err := store.Get().GetConfig(accountsField, &list)
	if err != nil && !errors.Is(err, store.ErrKeyNotFound) {
		return List{}, err
	}
	for i, account := range list.Items {
		if err = r.setAccountSecrets(&account); err != nil {
			return List{}, err
		}
		list.Items[i] = account
	}
	return list, nil
}

func (r defaultRepository) SaveAccounts(_ context.Context, list List) error {
	return store.Get().SetConfig(accountsField, list)
}

func (r defaultRepository) AddAccount(ctx context.Context, account Account) error {
	list, err := r.ListAccounts(ctx)
	if err != nil {
		return err
	}
	list.Add(account)
	if err = r.SaveAccounts(ctx, list); err != nil {
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

func (r defaultRepository) GetAccountById(ctx context.Context, id string) (Account, bool, error) {
	list, err := r.ListAccounts(ctx)
	if err != nil {
		return Account{}, false, err
	}
	if acc, ok := list.GetById(id); ok {
		if err = r.setAccountSecrets(&acc); err != nil {
			return Account{}, false, err
		}
		return acc, true, nil
	}
	return Account{}, false, nil
}

func (r defaultRepository) setAccountSecrets(acc *Account) (err error) {
	if acc.Id, err = store.Get().GetSecret(accountIdField); err != nil {
		return err
	}
	if acc.Token, err = store.Get().GetSecret(accountTokenField); err != nil {
		return err
	}
	return nil
}
