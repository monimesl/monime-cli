package cmds

import (
	"context"
	"errors"
	"github.com/monimesl/monime-cli/cli-utils/monimeapis"
	"github.com/monimesl/monime-cli/pkg/cmds/account"
	"github.com/monimesl/monime-cli/pkg/cmds/space"
	"github.com/monimesl/monime-cli/pkg/cmds/start"
	errors2 "github.com/monimesl/monime-cli/pkg/errors"
	"github.com/monimesl/monime-cli/pkg/text"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{
	Use:           "monime",
	Short:         "Monime command line tool for development and utility operations",
	Version:       "0.0.1",
	SilenceErrors: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand(
		space.Command,
		account.Command,
		start.Command,
	)
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
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
