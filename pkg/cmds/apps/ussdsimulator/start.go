package ussdsimulator

import (
	"github.com/monimesl/monime-cli/internal/resource/account"
	text2 "github.com/monimesl/monime-cli/internal/text"
	"github.com/monimesl/monime-cli/pkg/cobras"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var start = &cobra.Command{
	Use:   "start",
	Args:  cobras.NoArgs,
	Short: "Start the USSD Simulator",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := account.LoadActiveToken(cmd.Context()); err != nil {
			return err
		}
		text2.PrintStart("Starting the USSD simulator")
		ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer stop()
		_, err := launchApp(ctx)
		return err
	},
	SilenceUsage: true,
}
