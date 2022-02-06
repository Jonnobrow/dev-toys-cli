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

type Model struct {
	// Keymap and Help
	keys keyMap
	help help.Model
	// Result of last operation
	Result        string
	WriteToStdout bool
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
	stdin      string
}

func NewModel() Model {
	return Model{
		keys:     keys,
		help:     help.New(),
		title:    "dev-toys-cli",
		subtitle: "a swiss army knife for developers",
		categories: []*commands.Category{
			// Converters
			commands.NewCategory(
				"Converters", "Convert Stuff",
				commands.Converters()...,
			),
			// Formatters
			commands.NewCategory(
				"Formatters", "Format Stuff",
				commands.Formatters()...,
			),
		},
		inCategory: false,
	}
}

func (m Model) Init() tea.Cmd {
	return helpers.ReadFromStdin
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case helpers.StdinMsg:
		m.stdin = string(msg)
	case helpers.ErrorMsg:
		m.notification = fmt.Sprintf("Error: %v", msg.Error())
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
					m.Result = res
				}
			}
		case key.Matches(msg, m.keys.Toggle):
			m.clipboard = !m.clipboard
		case key.Matches(msg, m.keys.Yank):
			err := helpers.WriteToClipboard(m.Result)
			if err != nil {
				m.notification = "Failed to copy to clipboard"
			} else {
				m.notification = "Copied output to clipboard"
			}
		case key.Matches(msg, m.keys.Pipe):
			m.WriteToStdout = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
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

func (m *Model) goUp() {
	if m.inCategory {
		category := m.categories[m.cursor]
		category.CursorUp()
	} else {
		if m.cursor > 0 {
			m.cursor--
		}
	}
}

func (m *Model) goDown() {
	if m.inCategory {
		category := m.categories[m.cursor]
		category.CursorDown()
	} else {
		if m.cursor < len(m.categories)-1 {
			m.cursor++
		}
	}
}

func (m *Model) getInput() string {
	if m.clipboard {
		res, err := helpers.ReadFromClipboard()
		if err != nil {
			m.notification = err.Error()
			return ""
		}
		return res
	} else {
		return m.stdin
	}
}
