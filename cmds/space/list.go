package space

import (
	"context"
	"github.com/monimesl/monime-cli/internal/space"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Monimeer Spaces",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return listSpaces(cmd.Context())
	},
}

func listSpaces(ctx context.Context) error {
	svc, err := space.NewService()
	if err != nil {
		return err
	}
	return svc.ShowSpaceList(ctx)
}
