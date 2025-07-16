package cmds

import (
	"fmt"
	"github.com/monimesl/monime-cli/pkg/cobras"
	"github.com/monimesl/monime-cli/pkg/version"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	cmd *cobra.Command
}

func createVersionCommand() *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  cobras.NoArgs,
			Short: "Check the Monime CLI version",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(version.String)
			},
		},
	}
}
