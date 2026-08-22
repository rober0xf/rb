package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordEditCmd = &cobra.Command{
	Use:   "edit [entry]",
	Short: "Edit a specific entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("edit", args[0])
	},
}