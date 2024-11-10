package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	red   = lipgloss.Color("#9B2121")
	white = lipgloss.Color("#F1F1F1")

	bold = lipgloss.NewStyle().
		Bold(true)

	italic = lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("168"))

	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		Background(red).
		Width(80).
		Margin(0).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(white)

	commandBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(red).
			Width(80).
			Padding(1, 2)
)
