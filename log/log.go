package log

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	activeInputIndex int
	inputs           []textinput.Model
	returnToStart    func() tea.Model
}

type field struct {
	label       string
	placeholder string
}

var fields = [3]field{
	{
		label:       "Dose",
		placeholder: "1,5 mg",
	},
	{
		label:       "Substance",
		placeholder: "Estradiol",
	},
	{
		label:       "Route",
		placeholder: "Transdermal",
	},
}

const (
	DOSE_INDEX int = iota
	SUBSTANCE_INDEX
	ROUTE_INDEX
)

func InitializeModel(returnToStart func() tea.Model) tea.Model {
	var activeInputIndex int

	inputs := make([]textinput.Model, 3)

	for i := range inputs {
		inputs[i] = textinput.New()
		inputs[i].Width = 20
		inputs[i].Prompt = fmt.Sprintf("%s:\n", fields[i].label)
		//inputs[i].PromptStyle.Align(lipgloss.Center)
		inputs[i].Placeholder = fields[i].placeholder
	}

	inputs[0].Focus()

	return model{
		activeInputIndex,
		inputs,
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
		case "tab":
			if m.activeInputIndex >= 0 {
				m.inputs[m.activeInputIndex].Blur()
			}

			m.activeInputIndex++

			if m.activeInputIndex >= len(m.inputs) {
				m.activeInputIndex = 0
			}

			m.inputs[m.activeInputIndex].Focus()
		default:
			if m.activeInputIndex == -1 {
				switch msg.String() {
				case "q", "esc":
					return m.returnToStart(), nil
				}
			} else {
				switch msg.String() {
				case "esc":
					m.inputs[m.activeInputIndex].Blur()

					m.activeInputIndex = -1
				}
			}
		}
	}

	cmds := make([]tea.Cmd, len(m.inputs))

	for i := 0; i < len(m.inputs); i++ {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

var inputStyle = lipgloss.NewStyle().
	Padding(1).Margin(1).
	Border(lipgloss.ThickBorder()).
	Width(25).Align(lipgloss.Center)

func (m model) View() string {
	renderedFields := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		renderedFields[i] = inputStyle.Render(m.inputs[i].View())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedFields...)
}
