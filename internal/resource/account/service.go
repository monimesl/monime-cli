package account

import (
	"context"
	"fmt"
	"github.com/monime-lab/gwater"
	"github.com/monimesl/monime-cli/internal/resource/account/login"
	text2 "github.com/monimesl/monime-cli/internal/text"
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
		alias := text2.Format(acc.Alias, text2.FormatOptions{Bold: true, Color: "green"})
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
	if err = s.repository.AddAccount(ctx, account); err != nil {
		return err
	}
	alias := text2.Format(token.Account.Alias, text2.FormatOptions{Bold: true, Color: "green"})
	text2.PrintSuccess("Logged in successfully as user: %s", alias)
	return nil
}

func (s *Service) Logout(ctx context.Context, alias string) error {
	text2.PrintStart("Logging out of %s as user", alias)
	acc, exist, err := s.repository.GetAccountById(ctx, alias)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("unable to log out user %s as it does not exist", alias)
	}
	if err = s.repository.RemoveAccount(ctx, acc); err != nil {
		return err
	}
	text2.PrintSuccess("Logged out user %s successfully", alias)
	return nil
}
