package space

import (
	"context"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return activate(cmd.Context())
	},
}

func activate(ctx context.Context) error {
	return nil
}
