package log

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func increment(raw_before_delimiter *string, field *int) error {
	val, err := strconv.Atoi(*raw_before_delimiter)
	*raw_before_delimiter = ""

	if err == nil {
		*field += val
	}

	return err
}

func parseTime(raw string) (int, error) {
	days := 0
	hours := 0
	minutes := 0

	raw_before_delimiter := ""

	for i := range raw {
		r := raw[i]
		var err error

		switch r {
		case 'd':
			err = increment(&raw_before_delimiter, &days)
		case 'h':
			err = increment(&raw_before_delimiter, &hours)
		case 'm':
			err = increment(&raw_before_delimiter, &minutes)
		default:
			raw_before_delimiter += string(r)
		}

		if err != nil {
			return 0, err
		}
	}

	if len(raw_before_delimiter) > 0 {
		err := increment(&raw_before_delimiter, &minutes)

		if err != nil {
			return 0, err
		}
	}

	hours += days * 24
	minutes += hours * 60

	return minutes, nil
}

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
				m.input.Prompt = "Minutes since dose:\n(Must be an integer)\n"
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
