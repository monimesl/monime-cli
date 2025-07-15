package simulate

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "simulate",
	Short: "Simulate Webhook and 715 USSD short code",
}

func init() {
	Command.AddCommand(ussdSimulator, webhookSimulator)
}
