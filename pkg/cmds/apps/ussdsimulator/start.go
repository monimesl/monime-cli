package ussdsimulator

import (
	"errors"
	"os"
	"os/signal"

	"github.com/monimesl/monime-cli/internal/resource/account"
	text2 "github.com/monimesl/monime-cli/internal/text"
	"github.com/monimesl/monime-cli/pkg/cobras"
	"github.com/spf13/cobra"
)

var (
	exchangeURL    string
	exchangeMsisdn string
)

func newStartCommand() *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Args:  cobras.NoArgs,
		Short: "Start the USSD Simulator",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setGatewayEnvVars(); err != nil {
				return err
			}
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
	start.Flags().StringVar(
		&exchangeURL,
		"exchange-gateway-url",
		"",
		"The URL of the server handling the USSD session exchanges. (can also be set via MONIME_CLI_USSD_GATEWAY_EXCHANGE_URL)",
	)

	start.Flags().StringVar(
		&exchangeMsisdn,
		"exchange-msisdn",
		"",
		"The MSISDN to simulate in the USSD session. (Can also be set via MONIME_CLI_USSD_GATEWAY_EXCHANGE_MSISDN); Can only be set when 'exchange-gateway-url' is set.",
	)
	return start
}

func setGatewayEnvVars() (err error) {
	if exchangeURL != "" {
		err = errors.Join(err, os.Setenv("MONIME_CLI_USSD_GATEWAY_EXCHANGE_URL", exchangeURL))
		if exchangeMsisdn != "" {
			err = errors.Join(err, os.Setenv("MONIME_CLI_USSD_GATEWAY_EXCHANGE_MSISDN", exchangeMsisdn))
		}
	}
	return err
}
