package startMenu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type StartMenuModel struct {
	cursor int
}

type MenuItem struct {
	name     string
	getModel func() tea.Model
}

func InitializeModel() tea.Model {
	return StartMenuModel{
		cursor: 0,
	}
}

var menuItems = []MenuItem{
	{
		name:     "Log a dose",
		getModel: InitializeModel,
	},
	{
		name:     "View logs",
		getModel: InitializeModel,
	},
}

func (m StartMenuModel) Init() tea.Cmd {
	return nil
}

func (m StartMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor -= 1
				}
			case "down", "j":
				if m.cursor < len(menuItems)-1 {
					m.cursor += 1
				}
			case "enter":
				return menuItems[m.cursor].getModel(), nil
			}
		}
	}

	return m, nil
}

func (m StartMenuModel) View() string {
	s := "-= Skrive =-\n\n"

	for i, item := range menuItems {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item.name)
	}

	return s
}
