package webhookinspector

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "webhook-inspector",
	Short: "Start the Webhook Inspector",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("apps/webhook/build/bin/webhook.app/Contents/MacOS/webhook")
		return c.Run()
	},
}
