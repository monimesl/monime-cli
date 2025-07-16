package cmds

import (
	"context"
	"errors"
	"github.com/monimesl/monime-cli/cli-utils/monimeapis"
	errors2 "github.com/monimesl/monime-cli/internal/errors"
	"github.com/monimesl/monime-cli/internal/text"
	"github.com/monimesl/monime-cli/pkg/cmds/account"
	"github.com/monimesl/monime-cli/pkg/cmds/apps/ussdsimulator"
	"github.com/monimesl/monime-cli/pkg/cmds/apps/webhookinspector"
	"github.com/monimesl/monime-cli/pkg/cmds/space"
	"github.com/monimesl/monime-cli/pkg/version"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{
	Use:     "monime",
	Short:   "Monime command line tool for development and utility operations",
	Version: version.String,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(
		space.Command,
		account.Command,
		ussdsimulator.Command,
		webhookinspector.Command,
	)
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	rootCmd.Flags().BoolP("version", "v", false, "Check the Monime CLI version")
	rootCmd.AddCommand(createVersionCommand().cmd)
}

func ExecuteRootCmd() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	ctx2, cancel2 := context.WithTimeout(ctx, time.Minute*5)
	defer cancel2()
	if err := rootCmd.ExecuteContext(ctx2); err != nil {
		checkError(err)
	}
}

func checkError(err error) {
	msg := err.Error()
	switch {
	case errors.Is(err, context.Canceled):
		msg = "\nCommand cancelled"
	case errors.Is(err, context.DeadlineExceeded):
		msg = "\nCommand timed out"
	case errors.Is(err, errors2.ErrCliSilent):
		os.Exit(1)
	case errors.Is(err, monimeapis.ErrNotAuthenticated):
		errors2.PrintLoginHint()
		os.Exit(1)
	}
	text.PrintError(msg)
	os.Exit(1)
}
