package toys

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonnobrow/dev-toys-cli/internal/model"
)

func Run() error {
	program := tea.NewProgram(model.NewModel(
		map[string][]string{
			"encoding": {"base64", "url"},
			"decoding": {"base64", "url"},
		},
	))
	if err := program.Start(); err != nil {
		fmt.Printf("Oops, something went wrong: %v", err)
		return err
	}
	return nil
}
