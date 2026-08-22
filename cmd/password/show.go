package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordShowCmd = &cobra.Command{
	Use:   "show [entry]",
	Short: "Show the decrypted entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("show", args[0])
	},
}