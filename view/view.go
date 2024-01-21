package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"skrive.vanu.dev/logic"
)

type model struct {
	returnToStart func() tea.Model
	doses         []logic.Dose
	err           error

	loadingIndicator spinner.Model
}

func InitializeModel(returnToStart func() tea.Model) (tea.Model, tea.Cmd) {
	loadingIndicator := spinner.New()
	loadingIndicator.Spinner = spinner.Dot

	var doses []logic.Dose = nil
	var err error = nil

	model := model{
		returnToStart,
		doses,
		err,
		loadingIndicator,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(load, m.loadingIndicator.Tick)
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

	var cmd tea.Cmd
	m.loadingIndicator, cmd = m.loadingIndicator.Update(msg)

	return m, cmd
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
		return fmt.Sprintf("%s Loading", m.loadingIndicator.View())
	}
}
