package password

import (
	"github.com/spf13/cobra"
)

var passwordShowCmd = &cobra.Command{
	Use:   "show [entry]",
	Short: "Show the decrypted entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPass("show", args[0])
	},
}
