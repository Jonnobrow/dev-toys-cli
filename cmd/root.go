package cmd

import (
	"fmt"
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
			fmt.Fprintf(os.Stderr, "Error running: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
