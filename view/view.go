package view

import (
	"fmt"
	"strconv"

	"skrive/logic"
	"skrive/wrapper"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	returnToStart func() (tea.Model, tea.Cmd)
	doses         []logic.Dose
	err           error

	loadingIndicator spinner.Model
	doseTable        table.Model

	help help.Model
}

func InitializeModel(returnToStart func() (tea.Model, tea.Cmd)) (tea.Model, tea.Cmd) {
	loadingIndicator := spinner.New()
	loadingIndicator.Spinner = spinner.Dot

	doseTable := createTable(make([]logic.Dose, 0))

	var doses []logic.Dose = nil
	var err error = nil

	help := help.New()

	model := model{
		returnToStart,
		doses,
		err,
		loadingIndicator,
		doseTable,
		help,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(load, m.loadingIndicator.Tick)
}

func getTableRows(doses []logic.Dose) []table.Row {
	rows := make([]table.Row, len(doses))

	for i, dose := range doses {
		rows[i] = table.Row{
			strconv.Itoa(i + 1),
			dose.Time.Local().Format("2006-01-02 15:04:05"),
			dose.Quantity,
			dose.Substance,
			dose.Route,
		}
	}

	return rows
}

func createTable(doses []logic.Dose) table.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Time", Width: 20},
		{Title: "Amount", Width: 15},
		{Title: "Substance", Width: 25},
		{Title: "Route", Width: 25},
	}

	return table.New(
		table.WithColumns(columns),
		table.WithRows(getTableRows(doses)),
		table.WithFocused(true),
		table.WithHeight(10),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Exit):
			return m, tea.Quit
		case key.Matches(msg, keys.Quit):
			return m.returnToStart()
		case key.Matches(msg, keys.Delete):
			if len(m.doses) > 0 {
				index, _ := strconv.Atoi(
					m.doseTable.SelectedRow()[0],
				)

				return m, remove(m.doses, index-1)
			}
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	case successfulLoadMsg:
		m.doses = msg.doses
		m.doseTable = createTable(msg.doses)
	case failedLoadMsg:
		m.err = msg.err
	case removeMsg:
		m.doses = msg.doses
		m.doseTable.SetRows(getTableRows(m.doses))
		if m.doseTable.SelectedRow() == nil && len(m.doses) > 0 {
			m.doseTable.MoveUp(1)
		}
	}

	cmds := make([]tea.Cmd, 2)
	m.loadingIndicator, cmds[0] = m.loadingIndicator.Update(msg)
	m.doseTable, cmds[1] = m.doseTable.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var ui string

	if m.err != nil {
		ui = m.err.Error()
	} else if m.doses != nil {
		ui = fmt.Sprintf("%s\n%s",
			m.doseTable.View(),
			m.help.View(keys),
		)
	} else {
		ui = fmt.Sprintf("%s Loading", m.loadingIndicator.View())
	}

	return wrapper.Wrap(ui)
}
