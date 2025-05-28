package account

import (
	"slices"
	"time"
)

type Account struct {
	Id        string    `json:"-"`
	Token     string    `json:"-"`
	Alias     string    `json:"alias"`
	Active    bool      `json:"active"`
	Reference string    `json:"reference"`
	DateAdded time.Time `json:"date_added"`
}

type List struct {
	Items []Account `json:"items"`
}

func (l *List) GetActiveAccount() (Account, bool) {
	for _, a := range l.Items {
		if a.Active {
			return a, true
		}
	}
	return Account{}, false
}

func (l *List) GetById(idOrAlias string) (Account, bool) {
	idx := slices.IndexFunc(l.Items, func(e Account) bool {
		return e.Id == idOrAlias || e.Alias == idOrAlias
	})
	if idx == -1 {
		return Account{}, false
	}
	acc := l.Items[idx]
	return acc, true
}

func (l *List) Add(account Account) List {
	idx := slices.IndexFunc(l.Items, func(e Account) bool {
		return e.Reference == account.Reference
	})
	account.Active = true
	for _, a := range l.Items {
		a.Active = false
	}
	if idx == -1 {
		l.Items = append(l.Items, account)
	} else {
		l.Items[idx] = account
	}
	return *l
}

func (l *List) Remove(acc Account) {
	l.Items = slices.DeleteFunc(l.Items, func(a Account) bool {
		return a.Id == acc.Id
	})
}
