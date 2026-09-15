package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordCopyCmd = &cobra.Command{
	Use:   "copy [entry]",
	Short: "Copy a password to the clipboard",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("-c", args[0])
	},
}
