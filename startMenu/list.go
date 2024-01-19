package startMenu

import "github.com/charmbracelet/lipgloss"

var itemStyle = lipgloss.NewStyle()
var highlightedItemStyle = itemStyle.
	Background(lipgloss.Color("#D3D3D3")).
	Foreground(lipgloss.Color("#000000"))

var listStyle = lipgloss.NewStyle().Margin(1)

func renderListItem(value string, highlighted bool) string {
	if !highlighted {
		return itemStyle.Render(value)
	} else {
		return highlightedItemStyle.Render(value)
	}
}
