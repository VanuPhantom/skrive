package log

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	doseInput     textinput.Model
	returnToStart func() tea.Model
}

func InitializeModel(returnToStart func() tea.Model) tea.Model {
	doseInput := textinput.New()
	doseInput.Placeholder = "1,5 mg"
	doseInput.Focus()
	doseInput.Width = 20

	return model{
		doseInput,
		returnToStart,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			if m.doseInput.Focused() {
				switch msg.String() {
				case "esc":
					m.doseInput.Blur()
				}
			} else {
				switch msg.String() {
				case "q", "esc":
					return m.returnToStart(), nil
				}
			}
		}
	}

	var cmd tea.Cmd

	m.doseInput, cmd = m.doseInput.Update(msg)

	return m, cmd
}

var inputStyle = lipgloss.NewStyle().
	Padding(1).Margin(1).
	Border(lipgloss.ThickBorder()).
	Width(25).Align(lipgloss.Center)

func (m model) View() string {
	renderedDoseInput := inputStyle.Render("Dose:\n" + m.doseInput.View())
	renderedSubstanceInput := inputStyle.Render("Substance:\n" + m.doseInput.View())
	renderedRouteInput := inputStyle.Render("Route:\n" + m.doseInput.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedDoseInput, renderedSubstanceInput, renderedRouteInput)
}
