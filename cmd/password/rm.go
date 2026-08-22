package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordRemoveCmd = &cobra.Command{
	Use:   "remove [entry]",
	Short: "Remove a specific entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("rm", args[0])
	},
}