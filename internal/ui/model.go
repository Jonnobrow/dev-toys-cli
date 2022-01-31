package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jonnobrow/dev-toys-cli/internal/commands"
	"github.com/jonnobrow/dev-toys-cli/internal/helpers"
)

type model struct {
	// Keymap and Help
	keys keyMap
	help help.Model
	// Result of last operation
	result string
	// Notification
	notification string
	// Current view title
	title string
	// Current view subtitle
	subtitle string
	// categories
	categories []*commands.Category
	cursor     int
	width      int
	inCategory bool
	clipboard  bool
}

func NewModel() model {
	return model{
		keys:     keys,
		help:     help.New(),
		title:    "dev-toys-cli",
		subtitle: "a swiss army knife for developers",
		categories: []*commands.Category{
			// Converters
			commands.NewCategory(
				"Converters", "Convert Stuff", []commands.Command{
					commands.JSONtoYAML, commands.YAMLtoJSON,
				},
			),
			// ...
			commands.NewCategory(
				"Encoding", "Encode and Decode Stuff", []commands.Command{},
			),
		},
		inCategory: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.goUp()
		case key.Matches(msg, m.keys.Down):
			m.goDown()
		case key.Matches(msg, m.keys.Left):
			if m.inCategory {
				m.inCategory = false
			}
		case key.Matches(msg, m.keys.Right):
			fallthrough
		case key.Matches(msg, m.keys.Select):
			if !m.inCategory {
				m.inCategory = true
			} else {
				cat := m.categories[m.cursor]
				res, err := cat.Selected().Exec(m.getInput())
				if err != nil {
					m.notification = fmt.Sprintf("Error: %s", err.Error())
				} else {
					m.notification = "Success"
					m.result = res
				}
			}
		case key.Matches(msg, m.keys.Toggle):
			m.clipboard = !m.clipboard
		case key.Matches(msg, m.keys.Yank):
			err := helpers.WriteToClipboard(m.result)
			if err != nil {
				m.notification = "Failed to copy to clipboard"
			} else {
				m.notification = "Copied output to clipboard"
			}
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var sections []string

	sections = append(sections, m.headerView())

	if !m.inCategory {
		sections = append(sections, m.categoriesView())
	} else {
		sections = append(sections, m.subcommandsView())
	}

	sections = append(sections, m.help.View(m.keys))

	return appStyle.Render(lipgloss.JoinVertical(lipgloss.Left, sections...))
}

func (m *model) goUp() {
	if m.inCategory {
		category := m.categories[m.cursor]
		category.CursorUp()
	} else {
		if m.cursor > 0 {
			m.cursor--
		}
	}
}

func (m *model) goDown() {
	if m.inCategory {
		category := m.categories[m.cursor]
		category.CursorDown()
	} else {
		if m.cursor < len(m.categories)-1 {
			m.cursor++
		}
	}
}

func (m *model) getInput() string {
	if m.clipboard {
		res, err := helpers.ReadFromClipboard()
		if err != nil {
			m.notification = err.Error()
			return ""
		}
		return res
	} else {
		res, err := helpers.ReadFromStdin()
		if err != nil {
			m.notification = err.Error()
			return ""
		}
		return res
	}
}
