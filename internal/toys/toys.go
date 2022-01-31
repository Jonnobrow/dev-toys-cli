package toys

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonnobrow/dev-toys-cli/internal/ui"
)

func Run() error {
	program := tea.NewProgram(ui.NewModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if err := program.Start(); err != nil {
		fmt.Printf("Oops, something went wrong: %v", err)
		return err
	}
	return nil
}
