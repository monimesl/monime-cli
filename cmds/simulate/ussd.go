package simulate

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var ussdSimulator = &cobra.Command{
	Use:   "ussd",
	Short: "Open the USSD simulator",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("apps/ussd-simulator/build/bin/ussd-simulator.app/Contents/MacOS/ussd-simulator")
		return c.Run()
	},
}
