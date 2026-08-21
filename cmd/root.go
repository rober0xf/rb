package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rb/config"

	"github.com/spf13/cobra"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "rb",
	Short: "Personal CLI assistant",
}

func Execute(c *config.Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// save the config
	cfg = c

	return rootCmd.Execute()
}
