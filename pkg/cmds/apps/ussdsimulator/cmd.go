package ussdsimulator

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ussd-simulator",
	Short: "Manage the USSD Simulator App",
}

func init() {
	Command.AddCommand(start)
}
