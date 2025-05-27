package account

import (
	"context"
	"fmt"
	"github.com/monime-lab/gwater"
	"github.com/monimesl/monime-cli/internal/account/login"
	"github.com/monimesl/monime-cli/pkg/utils"
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

func (s *Service) ShowAccountList() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Account", "Active"})
	list, err := s.repository.ListAccounts()
	if err != nil {
		return err
	}
	for _, acc := range list.Items {
		alias := utils.Format(acc.Alias, utils.FormatOptions{Bold: true, Color: "green"})
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
	account.Alias = token.Account.Alias
	account.DateAdded = token.CreateTime
	account.Reference = gwater.UUID5(token.Account.Id)
	if err = s.repository.AddAccount(account); err != nil {
		return err
	}
	alias := utils.Format(token.Account.Alias, utils.FormatOptions{Bold: true, Color: "green"})
	fmt.Printf("✅ Logged in successfully as user: %s", alias)
	return nil
}
