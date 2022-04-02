package toys

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonnobrow/dev-toys-cli/internal/ui"
)

func Run() error {
	program := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if m, err := program.StartReturningModel(); err != nil {
		return err
	} else {
		if m, ok := m.(ui.Model); ok && m.WriteToStdout {
			fmt.Println(m.Result.Out())
		}
	}
	return nil
}
