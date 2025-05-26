package cmds

import (
	"github.com/monimesl/monime-cli/cmds/simulate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "monime",
	Short: "Monime command line tool for development and utility operations",
}

func init() {
	rootCmd.AddCommand(simulate.Command)

}

func ExecuteRootCmd() {
	cobra.CheckErr(rootCmd.Execute())
}
