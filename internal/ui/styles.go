package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	colors = []string{
		"#282828", "#cc241d", "#98971a", "#d79921", "#458558", "#b16286", "#689d6a", "#a89984",
		"#928374", "#fb4934", "#b8bb26", "#fabd2f", "#83a598", "#d3869b", "#8ec07c", "#ebdbb2",
	}

	// Get color profile
	profile = termenv.ColorProfile()

	// Base App Style
	appStyle = lipgloss.NewStyle().Padding(1, 0)

	headerStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			Foreground(lipgloss.AdaptiveColor{Dark: colors[13]})

	// Title Style
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingRight(1)
	subtitleStyle = lipgloss.NewStyle().
			PaddingRight(1)

	// List Styles
	listTitleStyle = lipgloss.NewStyle().
			Bold(true)

	listItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Dark: colors[3]}).
			Background(lipgloss.AdaptiveColor{Dark: colors[0]}).
			Margin(0, 0, 1, 1)

	listSelectedStyle = listItemStyle.Copy().
				Background(lipgloss.AdaptiveColor{Dark: colors[8]})
)
