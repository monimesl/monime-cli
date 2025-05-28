package account

import (
	"context"
	"github.com/monimesl/monime-cli/internal/account"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Monimeer accounts",
	Long: `Displays a list of all Monimeer accounts registered in the system.

This command retrieves and presents information such as Monimeer account ID, name, status, and creation time.
Useful for administrators or users managing multiple Monimeer accounts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return listAccounts(cmd.Context())
	},
}

func listAccounts(ctx context.Context) error {
	svc, err := account.NewService()
	if err != nil {
		return err
	}
	return svc.ShowAccountList(ctx)
}
