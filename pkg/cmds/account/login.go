package account

import (
	"context"
	"github.com/monimesl/monime-cli/pkg/resource/account"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Args:  cobra.NoArgs,
	Short: "Authenticate via browser and log into your Monime account",
	Long: `Starts a secure browser-based authentication flow.

This command opens a browser window for the user to authenticate with Monime. Upon successful login,
an access token is securely returned to the CLI. The token is then stored locally for future use.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return login(cmd.Context())
	},
}

func login(ctx context.Context) error {
	svc, err := account.NewService()
	if err != nil {
		return err
	}
	return svc.Login(ctx)
}
