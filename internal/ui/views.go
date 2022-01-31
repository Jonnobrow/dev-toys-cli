package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) titleView() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, titleStyle.Render(m.title), subtitleStyle.Render(m.subtitle))
}

func (m model) statusView() string {

	var statusComponents []string

	if m.notification != "" {
		statusComponents = append(statusComponents, fmt.Sprintf("✉ %s", m.notification))
	}

	if m.clipboard {
		statusComponents = append(statusComponents, fmt.Sprintf("Using clipboard as input"))
	} else {
		statusComponents = append(statusComponents, fmt.Sprintf("Using stdin as input"))
	}

	return strings.Join(statusComponents, " | ")
}

func (m model) headerView() string {
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

func (m model) categoriesView() string {

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

func (m model) subcommandsView() string {
	category := m.categories[m.cursor]
	var subcommands []string

	for i, command := range category.Subcommands {
		v := fmt.Sprintf("%s", listTitleStyle.Render(command.Name()))
		if i == category.Cursor {
			subcommands = append(subcommands, listSelectedStyle.Render(v))
		} else {
			subcommands = append(subcommands, listItemStyle.Render(v))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, subcommands...)
}
