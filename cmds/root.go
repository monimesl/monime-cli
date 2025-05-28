package cmds

import (
	"context"
	"errors"
	"github.com/monimesl/monime-cli/cmds/account"
	"github.com/monimesl/monime-cli/cmds/simulate"
	"github.com/monimesl/monime-cli/cmds/space"
	errors2 "github.com/monimesl/monime-cli/pkg/errors"
	"github.com/monimesl/monime-cli/pkg/utils/text"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{
	Use:           "monime",
	Short:         "Monime command line tool for development and utility operations",
	Version:       "0.1.2",
	SilenceErrors: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand(
		space.Command,
		account.Command,
		simulate.Command,
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
	case errors.Is(err, errors2.ErrAccountNotAuthenticated):
		errors2.PrintLoginHint()
		os.Exit(1)
	}
	text.PrintError(os.Stderr, msg)
	os.Exit(1)
}
