package log

import (
	"fmt"
	"skrive/wrapper"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	activeInputIndex int
	activeAreaIndex  int
	inputs           []textinput.Model
	popupModel       *popupModel
	returnToStart    func() (tea.Model, tea.Cmd)
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

var buttons = [2]string{
	"Log dose",
	"Backdate dose",
}

const (
	QUANTITY_INDEX int = iota
	SUBSTANCE_INDEX
	ROUTE_INDEX
)

func InitializeModel(returnToStart func() (tea.Model, tea.Cmd)) (tea.Model, tea.Cmd) {
	var (
		activeInputIndex int
		activeAreaIndex  int
		popupModel       *popupModel
	)

	inputs := make([]textinput.Model, 3)

	for i := range inputs {
		inputs[i] = textinput.New()
		inputs[i].Width = 20
		inputs[i].Prompt = fmt.Sprintf("%s:\n", fields[i].label)
		//inputs[i].PromptStyle.Align(lipgloss.Center)
		inputs[i].Placeholder = fields[i].placeholder
	}

	inputs[0].Focus()

	model := model{
		activeInputIndex,
		activeAreaIndex,
		inputs,
		popupModel,
		returnToStart,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) updateAfterFocusChange() (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 3)

	for i := range m.inputs {
		if m.activeAreaIndex == 0 && i == m.activeInputIndex {
			cmds[i] = m.inputs[i].Focus()
			continue
		}

		m.inputs[i].Blur()
	}

	return m, tea.Batch(cmds...)
}

func (m model) focusNext(allowEntry bool) (tea.Model, tea.Cmd) {
	if !allowEntry && m.activeInputIndex < 0 {
		return m, nil
	}

	if m.activeInputIndex == len(m.inputs)-1 {
		m.activeAreaIndex++

		if m.activeAreaIndex >= len(buttons)+1 {
			m.activeInputIndex = 0
			m.activeAreaIndex = 0
		}
	} else {
		m.activeInputIndex++
	}

	return m.updateAfterFocusChange()
}

func (m model) focusPrevious() (tea.Model, tea.Cmd) {
	if m.activeAreaIndex == 0 {
		m.activeInputIndex--

		if m.activeInputIndex < 0 {
			m.activeInputIndex = len(m.inputs) - 1
			m.activeAreaIndex = len(buttons)
		}
	} else {
		m.activeAreaIndex--
	}

	return m.updateAfterFocusChange()
}

func (m model) getValue(index int) string {
	return m.inputs[index].Value()
}

func (m model) Log(offset int) (tea.Model, tea.Cmd) {
	return m, log(
		m.getValue(QUANTITY_INDEX),
		m.getValue(SUBSTANCE_INDEX),
		m.getValue(ROUTE_INDEX),
		offset,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case logMsg:
		if msg.success {
			return m.returnToStart()
		} else {
			// TODO: Display errors
			return m, nil
		}
	default:
		if m.popupModel != nil {
			var (
				result *int
				cmd    tea.Cmd
			)

			result, m.popupModel, cmd = m.popupModel.Update(msg)

			if result != nil {
				return m.Log(*result)
			}

			if cmd != nil {
				return m, cmd
			}
		} else {
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "ctrl+c":
					return m, tea.Quit
				case "enter":
					if m.activeAreaIndex == 1 {
						return m.Log(0)
					} else if m.activeAreaIndex == 2 {
						popupModel := initializePopupModel()
						m.popupModel = &popupModel

						return m, m.popupModel.Init()
					} else {
						return m.focusNext(false)
					}
				case "tab":
					return m.focusNext(true)
				case "shift+tab":
					return m.focusPrevious()
				case "left":
					if m.activeAreaIndex == 0 {
						m.activeInputIndex--

						if m.activeInputIndex < 0 {
							m.activeInputIndex = len(m.inputs) - 1
						}

						return m.updateAfterFocusChange()
					}
				case "right":
					if m.activeAreaIndex == 0 {
						m.activeInputIndex++

						if m.activeInputIndex >= len(m.inputs) {
							m.activeInputIndex = 0
						}

						return m.updateAfterFocusChange()
					}
				case "down":
					if m.activeAreaIndex < len(buttons) {
						m.activeAreaIndex++
						return m.updateAfterFocusChange()
					}
				case "up":
					if m.activeAreaIndex > 0 {
						m.activeAreaIndex--
						return m.updateAfterFocusChange()
					}
				default:
					if m.activeInputIndex == -1 || m.activeAreaIndex > 0 {
						switch msg.String() {
						case "q", "esc":
							return m.returnToStart()
						}
					} else if m.activeAreaIndex == 0 {
						switch msg.String() {
						case "esc":
							m.activeInputIndex = -1
							return m.updateAfterFocusChange()
						}
					}
				}
			}
		}

		if m.popupModel == nil {

			cmds := make([]tea.Cmd, len(m.inputs))

			for i := 0; i < len(m.inputs); i++ {
				m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
			}

			return m, tea.Batch(cmds...)
		} else {
			return m, nil
		}
	}
}

var inputStyle = lipgloss.NewStyle().
	Padding(1).Margin(1).
	Border(lipgloss.ThickBorder()).
	Width(25).Align(lipgloss.Center)

var unhighlitButtonStyle = lipgloss.NewStyle().
	Padding(0, 1).Margin(0, 0, 1).
	Width(21)

var highlitButtonStyle = unhighlitButtonStyle.Copy().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFFFF"))

func buttonStyle(highlit bool) lipgloss.Style {
	if highlit {
		return highlitButtonStyle
	} else {
		return unhighlitButtonStyle
	}
}

func (m model) View() string {
	renderedFields := ""
	for i := 0; i < len(fields); i++ {
		renderedFields = lipgloss.JoinHorizontal(lipgloss.Top, renderedFields, inputStyle.Render(m.inputs[i].View()))
	}

	ui := lipgloss.JoinVertical(lipgloss.Center,
		renderedFields,
		buttonStyle(m.activeAreaIndex == 1).Render("Log dose"),
		buttonStyle(m.activeAreaIndex == 2).Render("Backdate dose"))

	if m.popupModel != nil {
		ui = m.popupModel.View(lipgloss.Size(ui))
	}

	return wrapper.Wrap(ui)
}
