package password

import (
	"github.com/spf13/cobra"
)

var passwordListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your passwords in a tree format",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runPass("ls")
	},
}
