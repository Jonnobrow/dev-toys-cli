package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-toys-cli",
	Short: "A swiss army knife for developers",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Welcome to DevToysCLI")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
