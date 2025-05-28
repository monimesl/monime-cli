package space

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "space",
	Short: "Manage Monime spaces",
}

func init() {
	Command.AddCommand(activateCmd, listCmd)
}
