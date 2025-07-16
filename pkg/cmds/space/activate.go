package space

import (
	"context"
	"fmt"
	errors2 "github.com/monimesl/monime-cli/internal/errors"
	"github.com/monimesl/monime-cli/internal/resource/space"
	"github.com/monimesl/monime-cli/internal/text"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate <space-id>",
	Short: "Activate a Space by its ID or alias.\n",
	Long: `This command sets the specified space as your current working context.
Subsequent operations (like managing webhooks or more) will be executed within the activated space.

You must provide a valid space ID or alias as an argument.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			text.PrintError("\033[1;31mMissing space ID or alias.\033[0m")
			fmt.Println("   Usage: monime space activate \033[1;33m<space-id>\033[0m")
			return errors2.ErrCliSilent
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return activate(cmd.Context(), args[0])
	},
}

func activate(ctx context.Context, idOrAlias string) error {
	svc, err := space.NewService()
	if err != nil {
		return err
	}
	return svc.ActivateSpace(ctx, idOrAlias)
}
