package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/jonnobrow/dev-toys-cli/internal/helpers"
	"github.com/spf13/cobra"
)

type Command interface {
	Name() string
	Exec(string) (string, error)
	DisplayInput(string) string
	DisplayOutput(string) string
	ShouldDisplayInput() bool
	CobraCommand(func(string) (string, error)) cobra.Command

	toCliName() string
}

type Result struct {
	commandName string
	result      string
}

func (r Result) IsFromCommand(cmd Command) bool {
	return r.commandName == cmd.Name()
}

func (r Result) Command() string {
	return r.commandName
}

func (r Result) Out() string {
	return r.result
}

func NewResult(result string, commandName string) Result {
	return Result{commandName, result}
}

type base struct {
	name string
	desc string

	displayInput  bool
	displayOutput bool

	cliName string
	aliases []string
}

func (b base) Name() string {
	return b.name
}

func (b base) toCliName() string {
	if b.cliName != "" {
		return b.cliName
	}
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(b.name, " ", "-"), "->", ""), "--", "-"))
}

func (b base) Desc() string {
	return b.desc
}

func (b base) DisplayInput(input string) string {
	return input
}

func (b base) DisplayOutput(input string) string {
	return input
}

func (b base) ShouldDisplayInput() bool {
	return b.displayInput
}

func (b base) Exec(string) (string, error) { return "", nil }

func NewBase(name, desc string) base {
	newBase := base{
		name:          name,
		desc:          desc,
		displayInput:  true,
		displayOutput: true,
	}

	return newBase
}

func (b base) withoutInputDisplay() base {
	var newB base
	newB = b
	newB.displayInput = false
	return newB
}

func (b base) withCliName(cliName string) base {
	var newB base
	newB = b
	newB.cliName = cliName
	return newB
}

func (b base) withAliases(aliases []string) base {
	var newB base
	newB = b
	newB.aliases = aliases
	return newB
}

func (b base) CobraCommand(Exec func(string) (string, error)) cobra.Command {
	return cobra.Command{
		Use:     strings.ToLower(b.toCliName()),
		Short:   b.desc,
		Aliases: b.aliases,
		Run: func(cmd *cobra.Command, args []string) {
			var input string
			hasStdin, err := helpers.HasStdin()
			if hasStdin {
				if stdinData, ok := helpers.ReadFromStdin().(helpers.StdinMsg); ok {
					input = string(stdinData)
				}
			} else {
				input = strings.Join(args, " ")
			}
			res, err := Exec(input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error running: %v", err)
				os.Exit(1)
			}
			fmt.Printf("%s", res)
		},
	}
}
