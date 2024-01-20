package startMenu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"skrive.vanu.dev/about"
	"skrive.vanu.dev/log"
)

type StartMenuModel struct {
	cursor int
}

type MenuItem struct {
	name     string
	getModel func(func() tea.Model) tea.Model
}

func InitializeModel() tea.Model {
	return StartMenuModel{
		cursor: 0,
	}
}

func getModelPlaceHolder(returnToStart func() tea.Model) tea.Model { return returnToStart() }

var menuItems = []MenuItem{
	{
		name:     "Log a dose",
		getModel: log.InitializeModel,
	},
	{
		name:     "View logs",
		getModel: getModelPlaceHolder,
	},
	{
		name:     "About",
		getModel: about.InitializeModel,
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
		case "up", "k":
			if m.cursor > 0 {
				m.cursor -= 1
			}
		case "down", "j":
			if m.cursor < len(menuItems)-1 {
				m.cursor += 1
			}
		case "enter":
			return menuItems[m.cursor].getModel(InitializeModel), nil
		}
	}

	return m, nil
}

func (m StartMenuModel) View() string {
	header := renderHeader()
	list := ""

	for i, item := range menuItems {
		if i > 0 {
			list += "\n"
		}
		list += renderListItem(item.name, m.cursor == i)
	}

	list = listStyle.Render(list)

	return lipgloss.JoinHorizontal(lipgloss.Center, header, list)
}
