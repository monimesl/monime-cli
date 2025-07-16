package cobras

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func NoArgs(cmd *cobra.Command, args []string) error {
	path := cmd.CommandPath()
	message := fmt.Sprintf("'%s' takes no positional arguments. Please run '%s --help' for usage information.", path, path)
	if len(args) > 0 {
		return errors.New(message)
	}
	return nil
}
