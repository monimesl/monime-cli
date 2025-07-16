package ussd

import (
	"fmt"
	"github.com/monimesl/monime-cli/pkg/resource/account"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var Simulator = &cobra.Command{
	Use:   "ussd",
	Short: "Open the USSD simulator",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := account.LoadActiveToken(cmd.Context()); err != nil {
			return err
		}
		fmt.Println("🚀 Opening the USSD simulator")
		ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer stop()
		_, err := launchApp(ctx)
		return err
	},
	SilenceUsage: true,
}
