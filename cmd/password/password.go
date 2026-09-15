package password

import (
	ptui "rb/internal/password"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	PasswordCommand.AddCommand(
		passwordListCmd,
		passwordShowCmd,
		passwordRemoveCmd,
		passwordEditCmd,
		passwordCreateCmd,
		passwordInitCmd,
		passwordCopyCmd,
	)
}

var PasswordCommand = &cobra.Command{
	Use:   "password",
	Short: "Manage your passwords",
}

var TUICommand = &cobra.Command{
	Use:   "ptui",
	Short: "Run the TUI interface",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := ptui.NewPasswordManager()
		if err != nil {
			return err
		}

		model, err := ptui.NewModel(manager)
		if err != nil {
			return err
		}

		_, err = tea.NewProgram(model).Run()
		return err
	},
}
