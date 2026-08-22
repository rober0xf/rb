package cmd

import (
	"github.com/spf13/cobra"
	"rb/internal/ai"
)

var aiCmd = &cobra.Command{
	Use:   "ai [input]",
	Short: "Interact with the model via CLI",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return ai.CallModel(cmd.Context(), cfg, args[0])
	},
}