package space

import (
	"context"
	"github.com/monimesl/monime-cli/internal/space"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idOrAlias := args[0]
		return activate(cmd.Context(), idOrAlias)
	},
}

func activate(ctx context.Context, idOrAlias string) error {
	svc, err := space.NewService()
	if err != nil {
		return err
	}
	return svc.ActivateSpace(ctx, idOrAlias)
}
