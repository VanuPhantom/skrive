package about

import (
	"skrive/wrapper"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var red = lipgloss.NewStyle().Background(lipgloss.Color("#C8102E"))
var white = lipgloss.NewStyle().Background(lipgloss.Color("#FFFFFF"))

type model struct {
	returnToStart func() (tea.Model, tea.Cmd)
}

func InitializeModel(returnToStart func() (tea.Model, tea.Cmd)) (tea.Model, tea.Cmd) {
	model := model{
		returnToStart,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			return m.returnToStart()
		}
	}

	return m, nil
}

var flagStyle = lipgloss.NewStyle().Margin(1)

func (m model) View() string {
	leftRed := red.Render("      \n      \n      ")
	rightRed := red.Render("          \n          \n          ")
	leftWhite := white.Render("      ")
	centerWhite := white.Render("  \n  \n  \n  \n  \n  \n  ")
	rightWhite := white.Render("          ")

	left := lipgloss.JoinVertical(lipgloss.Left, leftRed, leftWhite, leftRed)
	right := lipgloss.JoinVertical(lipgloss.Left, rightRed, rightWhite, rightRed)

	flag := flagStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, left, centerWhite, right))

	return wrapper.Wrap(
		lipgloss.JoinHorizontal(lipgloss.Center, flag, "Made in Denmark\nTess Ellenoir Duursma\nvanu.dev"),
	)
}
