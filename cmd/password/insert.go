package password

import (
	"github.com/spf13/cobra"
)

var passwordCreateCmd = &cobra.Command{
	Use:   "create [entry]",
	Short: "Create an entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPass("insert", args[0])
	},
}
