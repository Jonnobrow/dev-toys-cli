package cmd

import (
	"fmt"
	"os"

	"github.com/jonnobrow/dev-toys-cli/internal/commands"
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

func init() {
	for _, cat := range commands.Categories {
		catCommand := cobra.Command{
			Use:     cat.CliName(),
			Aliases: []string{string(cat.CliShort())},
			Short:   cat.Title,
			Long:    cat.Prompt,
		}
		for _, c := range cat.Subcommands {
			cmd := c.CobraCommand(c.Exec)
			catCommand.AddCommand(
				&cmd,
			)
		}
		rootCmd.AddCommand(&catCommand)
	}

}
