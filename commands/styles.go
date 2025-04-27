package commands

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	red   = lipgloss.Color("#9B2121")
	white = lipgloss.Color("#F1F1F1")
	pink  = lipgloss.Color("168")

	baseStyle = lipgloss.NewStyle().MarginLeft(1)

	bold = baseStyle.
		Bold(true).
		Foreground(pink)

	italic = baseStyle.
		Italic(true)

	ErrorMsg = baseStyle.
			Foreground(pink).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			BorderForeground(red)

	header = baseStyle.
		Bold(true).
		Foreground(white).
		Background(red).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(white)

	commandBox = baseStyle.
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(red).
			Padding(1, 2)
)

func printHeader(s string) {
	fmt.Println(header.Render(s))
}
