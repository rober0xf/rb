package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open [site]",
	Short: "Open a site in the browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return openBrowser(args[0])
	},
}

func openBrowser(name string) error {
	urls := map[string]string{
		"chatgpt": "https://chatgpt.com",
		"github":  "https://github.com/rober0xf",
		"reddit":  "https://reddit.com",
		"claude":  "https://claude.ai/new",
		"youtube": "https://youtube.com",
		"mail":    "https://mail.google.com/mail/u/0/",
	}

	url, ok := urls[name]
	if !ok {
		return nil
	}

	return exec.Command("xdg-open", url).Start()
}
