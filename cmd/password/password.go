package password

import (
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(
		passwordListCmd,
		passwordShowCmd,
		passwordRemoveCmd,
		passwordEditCmd,
		passwordCreateCmd,
	)
}

var Command = &cobra.Command{
	Use:   "password",
	Short: "Manage your passwords",
}
