package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	red   = lipgloss.Color("#9B2121")
	white = lipgloss.Color("#F1F1F1")

	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		Background(red).
		Width(60).
		Margin(0).
		Padding(1, 4).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(white)

	commandBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(red).
			Width(60).
			Padding(1, 4)
)
