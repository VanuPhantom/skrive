package wrapper

import "github.com/charmbracelet/lipgloss"

var wrapperStyle = lipgloss.NewStyle().MarginBottom(1)

func Wrap(ui string) string {
	return wrapperStyle.Render(ui)
}
