package password

import (
	"github.com/spf13/cobra"
)

var passwordRemoveCmd = &cobra.Command{
	Use:   "remove [entry]",
	Short: "Remove a specific entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPass("rm", args[0])
	},
}
