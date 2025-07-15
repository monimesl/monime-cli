package account

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "account",
	Short: "Manage Monime accounts",
}

func init() {
	Command.AddCommand(listCmd, loginCmd, logoutCmd)
}
