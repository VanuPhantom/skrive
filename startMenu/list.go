package startMenu

import "github.com/charmbracelet/lipgloss"

var itemStyle = lipgloss.NewStyle()
var highlightedItemStyle = itemStyle.
	Foreground(lipgloss.AdaptiveColor{Dark: "#000000", Light: "#FFFFFF"}).
	Background(lipgloss.AdaptiveColor{Dark: "#FFFFFF", Light: "#000000"})

var listStyle = lipgloss.NewStyle().Margin(1)

func renderListItem(value string, highlighted bool) string {
	if !highlighted {
		return itemStyle.Render(value)
	} else {
		return highlightedItemStyle.Render(value)
	}
}
