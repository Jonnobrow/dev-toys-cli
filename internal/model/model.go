package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	cursor      int
	categories  []string
	subcommands map[string][]string // Map of categories to commands
}

func NewModel(subcommands map[string][]string) (m Model) {
	var categories []string

	for cat := range subcommands {
		categories = append(categories, cat)
	}

	return Model{
		cursor:      0,
		categories:  categories,
		subcommands: subcommands,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.categories)-1 {
				m.cursor++
			}

		}

	}
	return m, nil
}

func (m Model) View() string {
	s := "What tool do you need today?\n\n"
	for i, category := range m.categories {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s : [%s]\n", cursor, category, strings.Join(m.subcommands[category], ", "))
	}
	return s
}
