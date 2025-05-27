package account

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "account",
	Short: "Manage Monimeer accounts",
}

func init() {
	Command.AddCommand(accountLoginCmd, accountListCmd)
}
