package cmd

import (
	"os"

	"github.com/jonnobrow/dev-toys-cli/internal/toys"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-toys-cli",
	Short: "A swiss army knife for developers",
	Run: func(cmd *cobra.Command, _ []string) {
		err := toys.Run()
		if err != nil {
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
