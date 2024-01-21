package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"skrive.vanu.dev/logic"
)

type model struct {
	returnToStart func() tea.Model
	doses         []logic.Dose
	err           error
}

func InitializeModel(returnToStart func() tea.Model) (tea.Model, tea.Cmd) {
	var doses []logic.Dose = nil
	var err error = nil

	model := model{
		returnToStart,
		doses,
		err,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return load
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			return m.returnToStart(), nil
		}
	case successfulLoadMsg:
		m.doses = msg.doses
	case failedLoadMsg:
		m.err = msg.err
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	} else if m.doses != nil {
		result := ""

		for _, dose := range m.doses {
			result += fmt.Sprintf("- %s - %s - %s\n", dose.Quantity, dose.Substance, dose.Route)
		}

		return result
	} else {
		return "Loading"
	}
}
