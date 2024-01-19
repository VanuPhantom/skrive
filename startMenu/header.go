package startMenu

import "github.com/charmbracelet/lipgloss"

var blueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5BCEFA"))
var pinkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F5A9B8"))
var whiteStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))

var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Border(lipgloss.ThickBorder()).
	Padding(1, 2)

func renderHeader() string {
	s := ""

	s += blueStyle.Render("S")
	s += pinkStyle.Render("k")
	s += whiteStyle.Render("r")
	s += whiteStyle.Render("i")
	s += pinkStyle.Render("v")
	s += blueStyle.Render("e")

	return headerStyle.Render(s)
}
