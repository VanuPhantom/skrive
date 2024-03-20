package log

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type popupModel struct {
	input textinput.Model
}

func initializePopupModel() popupModel {
	input := textinput.New()
	input.Prompt = "Minutes since dose:\n"
	input.Focus()

	return popupModel{
		input,
	}
}

func (m popupModel) Init() tea.Cmd {
	return tea.Batch(m.input.Focus(), textinput.Blink)
}

func (m popupModel) Update(msg tea.Msg) (*int, *popupModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return nil, &m, tea.Quit
		case "enter":
			raw_value := m.input.Value()
			value, err := parseTime(raw_value)

			if err != nil {
				m.input.Prompt = "Minutes since dose:\n(1d2h3m or 123 (for 123 minutes))\n"
			} else {
				return &value, nil, nil
			}
		case "esc":
			return nil, nil, nil
		}
	}

	var cmd tea.Cmd

	m.input, cmd = m.input.Update(msg)

	return nil, &m, cmd
}

var popupStyle = lipgloss.NewStyle().
	Padding(1).
	Width(35).Height(5).
	Align(lipgloss.Center, lipgloss.Center).
	Border(lipgloss.ThickBorder())

func (m popupModel) View(containerWidth int, containerHeight int) string {
	return lipgloss.Place(containerWidth, containerHeight,
		lipgloss.Center, lipgloss.Center,
		popupStyle.Render(m.input.View()),
		lipgloss.WithWhitespaceChars("@#"),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#3A3B3C")),
	)
}
