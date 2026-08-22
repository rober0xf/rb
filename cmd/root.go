package cmd

import (
	"context"
	"os"
	"os/signal"
	"rb/cmd/password"
	"rb/config"

	"github.com/spf13/cobra"
)

var cfg *config.Config

func init() {
	rootCmd.AddCommand(
		password.Command,
		openCmd,
		aiCmd,
	)
}

var rootCmd = &cobra.Command{
	Use:   "rb",
	Short: "Personal CLI assistant",
}

func Execute(c *config.Config) error {
	// save the config
	cfg = c

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	return rootCmd.ExecuteContext(ctx)
}
