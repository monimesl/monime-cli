package webhookinspector

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "webhook-inspector",
	Short: "Manage the Webhook Inspector App",
}

func init() {
	Command.AddCommand(start)
}
