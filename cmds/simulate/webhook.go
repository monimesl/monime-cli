package simulate

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var webhookSimulator = &cobra.Command{
	Use:   "webhook",
	Short: "Open the Webhook simulator",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("open /Users/alphashaw/projects/Projects/sources/monimesl/apps/monime-cli/apps/webhook/build/bin/webhook.app")
		return c.Run()
	},
}
