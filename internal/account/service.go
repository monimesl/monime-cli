package account

import (
	"context"
	"fmt"
	"github.com/monime-lab/gwater"
	"github.com/monimesl/monime-cli/internal/account/login"
	"github.com/monimesl/monime-cli/pkg/utils/text"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func NewService() (*Service, error) {
	return &Service{
		repository: &defaultRepository{},
	}, nil
}

type Service struct {
	loginFLow  login.Flow
	repository Repository
}

func (s *Service) GetActiveAccount(ctx context.Context) (Account, bool, error) {
	list, err := s.repository.ListAccounts(ctx)
	if err != nil {
		return Account{}, false, err
	}
	acc, ok := list.GetActiveAccount()
	return acc, ok, nil
}

func (s *Service) ShowAccountList(ctx context.Context) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Account", "Active"})
	list, err := s.repository.ListAccounts(ctx)
	if err != nil {
		return err
	}
	for _, acc := range list.Items {
		alias := text.Format(acc.Alias, text.FormatOptions{Bold: true, Color: "green"})
		if err = table.Append([]string{alias, strconv.FormatBool(acc.Active)}); err != nil {
			return err
		}
	}
	return table.Render()
}

func (s *Service) Login(ctx context.Context) error {
	token, err := s.loginFLow.Run(ctx)
	if err != nil {
		return err
	}
	account := Account{}
	account.Id = token.Account.Id
	account.Token = token.Id
	account.Alias = token.Account.Alias
	account.DateAdded = token.CreateTime
	account.Reference = gwater.UUID5(token.Account.Id)
	if err = s.repository.AddAccount(account); err != nil {
		return err
	}
	alias := text.Format(token.Account.Alias, text.FormatOptions{Bold: true, Color: "green"})
	fmt.Printf("✅ Logged in successfully as user: %s", alias)
	return nil
}
