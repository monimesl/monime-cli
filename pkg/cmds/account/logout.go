package account

import (
	"context"
	errors2 "github.com/monimesl/monime-cli/pkg/errors"
	"github.com/monimesl/monime-cli/pkg/resource/account"
	"github.com/monimesl/monime-cli/pkg/text"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout <alias>",
	Short: "Log out of your Monime account",
	Long: `Log out of your Monime account by clearing locally stored credentials.

This command deletes your authentication token from local storage.
You will need to run 'monime account login' again to reauthenticate.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			text.PrintError("\033[1;31mMissing account alias\033[0m")
			return errors2.ErrCliSilent
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return logout(cmd.Context(), args[0])
	},
}

func logout(ctx context.Context, alias string) error {
	svc, err := account.NewService()
	if err != nil {
		return err
	}
	return svc.Logout(ctx, alias)
}
