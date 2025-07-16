package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/monimesl/monime-cli/internal/store"
)

type Repository interface {
	ListAccounts(context.Context) (List, error)
	SaveAccounts(ctx context.Context, list List) error
	AddAccount(ctx context.Context, account Account) error
	GetAccountById(ctx context.Context, id string) (Account, bool, error)
	RemoveAccount(ctx context.Context, acc Account) error
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
		if err = r.loadAccountSecrets(&account); err != nil {
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
	if err = r.storeAccountSecrets(account); err != nil {
		return err
	}
	return nil
}

func (r defaultRepository) RemoveAccount(ctx context.Context, acc Account) error {
	list, err := r.ListAccounts(ctx)
	if err != nil {
		return err
	}
	list.Remove(acc)
	if err = r.SaveAccounts(ctx, list); err != nil {
		return err
	}
	if err = r.deleteAccountSecrets(acc); err != nil {
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
		if err = r.loadAccountSecrets(&acc); err != nil {
			return Account{}, false, err
		}
		return acc, true, nil
	}
	return Account{}, false, nil
}

func (r defaultRepository) loadAccountSecrets(acc *Account) (err error) {
	idKey := r.accountIdSecretKey(*acc)
	tokenKey := r.accountTokenSecretKey(*acc)
	if acc.Id, err = store.Get().GetSecret(idKey); err != nil && !errors.Is(err, store.ErrKeyNotFound) {
		return err
	}
	if acc.Token, err = store.Get().GetSecret(tokenKey); err != nil && !errors.Is(err, store.ErrKeyNotFound) {
		return err
	}
	return nil
}

func (r defaultRepository) storeAccountSecrets(acc Account) (err error) {
	idKey := r.accountIdSecretKey(acc)
	tokenKey := r.accountTokenSecretKey(acc)
	if err = store.Get().SetSecret(idKey, acc.Id); err != nil {
		return err
	}
	if err = store.Get().SetSecret(tokenKey, acc.Token); err != nil {
		return err
	}
	return nil
}

func (r defaultRepository) deleteAccountSecrets(acc Account) (err error) {
	return errors.Join(
		store.Get().DeleteSecret(r.accountIdSecretKey(acc)),
		store.Get().DeleteSecret(r.accountTokenSecretKey(acc)),
	)
}

func (r defaultRepository) accountIdSecretKey(acc Account) string {
	return fmt.Sprintf("%s_%s", accountIdField, acc.Reference)
}

func (r defaultRepository) accountTokenSecretKey(acc Account) string {
	return fmt.Sprintf("%s_%s", accountTokenField, acc.Reference)
}
