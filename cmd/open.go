package cmd

import (
	"github.com/spf13/cobra"
	"rb/internal/browser"
)

var openCmd = &cobra.Command{
	Use:   "open [site]",
	Short: "Open a site in the browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return browser.OpenBrowser(args[0])
	},
}
