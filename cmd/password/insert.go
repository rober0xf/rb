package password

import (
	"github.com/spf13/cobra"
	"rb/internal/password"
)

var passwordCreateCmd = &cobra.Command{
	Use:   "create [entry]",
	Short: "Create an entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return password.RunPass("insert", args[0])
	},
}