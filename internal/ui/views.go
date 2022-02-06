package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wrap"
)

func (m Model) titleView() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, titleStyle.Render(m.title), subtitleStyle.Render(m.subtitle))
}

func (m Model) statusView() string {

	var statusComponents []string

	if m.notification != "" {
		statusComponents = append(statusComponents, fmt.Sprintf("✉ %s", m.notification))
	}

	if m.clipboard {
		statusComponents = append(statusComponents, fmt.Sprintf("Using clipboard as input"))
	} else {
		statusComponents = append(statusComponents, fmt.Sprintf("Using stdin as input"))
	}

	if m.Result.Command() != "" && m.Result.Out() != "" {
		statusComponents = append(statusComponents, fmt.Sprintf("Last command: %s", m.Result.Command()))
	}

	return strings.Join(statusComponents, " | ")
}

func (m Model) headerView() string {
	availWidth := m.width
	title := m.titleView()
	availWidth -= lipgloss.Width(title)

	status := m.statusView()
	spacerWidth := availWidth - lipgloss.Width(status)

	var v string
	if spacerWidth > 0 {
		v = lipgloss.JoinHorizontal(lipgloss.Left, title, lipgloss.NewStyle().Width(spacerWidth).Render(""), status)
	} else {
		v = lipgloss.JoinVertical(lipgloss.Left, title, status)
	}

	return headerStyle.Render(
		lipgloss.PlaceHorizontal(m.width, lipgloss.Left, v))
}

func (m Model) categoriesView() string {

	var categories []string

	for i, category := range m.categories {
		v := fmt.Sprintf("%s\n%s", listTitleStyle.Render(category.Title), category.Prompt)
		if i == m.cursor {
			categories = append(categories, listSelectedStyle.Render(v))
		} else {
			categories = append(categories, listItemStyle.Render(v))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, categories...)
}

func (m Model) subcommandsView() string {
	category := m.categories[m.cursor]
	var subcommands []string

	for i, command := range category.Subcommands {
		v := listTitleStyle.Render(command.Name())
		if i == category.Cursor {
			if category.Selected().ShouldDisplayInput() {
				v = fmt.Sprintf("%s\n%s\n", v, m.inputOutputView())
			} else {
				v = fmt.Sprintf("%s\n%s\n", v, m.outputView())
			}
			subcommands = append(subcommands, listSelectedStyle.Render(v))
		} else {
			subcommands = append(subcommands, listItemStyle.Render(v))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, subcommands...)
}

func (m Model) inputOutputView() string {
	category := m.categories[m.cursor]
	subcommand := category.Selected()
	if !m.Result.IsFromCommand(subcommand) {
		return "Select to Run Command"
	}

	delimeter := lipgloss.NewStyle().Bold(true).Render(" to ")
	availWidth := m.width - lipgloss.Width(delimeter)

	input := lipgloss.NewStyle().MaxHeight(20).Render(wrap.String(subcommand.DisplayInput(m.getInput()), availWidth/2))
	output := lipgloss.NewStyle().MaxHeight(20).Render(wrap.String(subcommand.DisplayOutput(m.Result.Out()), availWidth/2))
	return lipgloss.JoinHorizontal(lipgloss.Left, input, " to ", output)
}

func (m Model) outputView() string {
	category := m.categories[m.cursor]
	subcommand := category.Selected()
	if !m.Result.IsFromCommand(subcommand) {
		return "Select to Run Command"
	}
	result, err := subcommand.Exec(m.getInput())
	if err != nil {
		return "Error running command"
	}
	return lipgloss.NewStyle().MaxHeight(20).Render(wrap.String(subcommand.DisplayOutput(result), m.width))
}
