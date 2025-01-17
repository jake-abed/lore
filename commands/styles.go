package commands

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	red   = lipgloss.Color("#9B2121")
	white = lipgloss.Color("#F1F1F1")
	pink  = lipgloss.Color("168")

	bold = lipgloss.NewStyle().
		Bold(true).
		Foreground(pink)

	italic = lipgloss.NewStyle().
		Italic(true)

	ErrorMsg = lipgloss.NewStyle().
			Foreground(pink).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			BorderForeground(red)

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

func printHeader(s string) {
	fmt.Println(header.Render(s))
}
