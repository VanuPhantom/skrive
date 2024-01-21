package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"skrive.vanu.dev/logic"
)

type model struct {
	returnToStart func() tea.Model
	doses         []logic.Dose
	err           error

	loadingIndicator spinner.Model
	doseTable        table.Model
}

func InitializeModel(returnToStart func() tea.Model) (tea.Model, tea.Cmd) {
	loadingIndicator := spinner.New()
	loadingIndicator.Spinner = spinner.Dot

	doseTable := createTable(make([]logic.Dose, 0))

	var doses []logic.Dose = nil
	var err error = nil

	model := model{
		returnToStart,
		doses,
		err,
		loadingIndicator,
		doseTable,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(load, m.loadingIndicator.Tick)
}

func createTable(doses []logic.Dose) table.Model {
	columns := []table.Column{
		{Title: "Amount", Width: 15},
		{Title: "Substance", Width: 30},
		{Title: "Route", Width: 30},
	}

	rows := make([]table.Row, len(doses))

	for i, dose := range doses {
		rows[i] = table.Row{
			dose.Quantity,
			dose.Substance,
			dose.Route,
		}
	}

	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
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
		m.doseTable = createTable(msg.doses)
	case failedLoadMsg:
		m.err = msg.err
	}

	cmds := make([]tea.Cmd, 2)
	m.loadingIndicator, cmds[0] = m.loadingIndicator.Update(msg)
	m.doseTable, cmds[1] = m.doseTable.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	} else if m.doses != nil {
		return m.doseTable.View()
	} else {
		return fmt.Sprintf("%s Loading", m.loadingIndicator.View())
	}
}
