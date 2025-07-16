package start

import (
	"github.com/monimesl/monime-cli/pkg/cmds/start/apps/ussd"
	"github.com/monimesl/monime-cli/pkg/cmds/start/apps/webhook"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "start",
	Short: "Start a CLI app (USSD Simulator, Webhook Inspector)",
}

func init() {
	Command.AddCommand(ussd.Simulator, webhook.Inspector)
}
