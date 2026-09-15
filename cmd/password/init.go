package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordInitCmd = &cobra.Command{
	Use:   "init [gpg-id]",
	Short: "Initialize the password store",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("init", args[0])
	},
}
