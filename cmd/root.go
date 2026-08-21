package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rb",
	Short: "Personal CLI assistant",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
