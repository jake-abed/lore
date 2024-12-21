package commands

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	red   = lipgloss.Color("#9B2121")
	white = lipgloss.Color("#F1F1F1")

	bold = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("168"))

	italic = lipgloss.NewStyle().
		Italic(true)

	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		Background(red).
		Margin(0).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(white)

	commandBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(red).
			Padding(1, 2)
)
