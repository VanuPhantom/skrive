package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type StartMenuModel struct {
	cursor int
}

var menuItems = []string{"Log a dose", "View logs"}

func initializeStartMenuModel() StartMenuModel {
	return StartMenuModel{
		cursor: 0,
	}
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
			}
		}
	}

	return m, nil
}

func (m StartMenuModel) View() string {
	s := "-= Skrive =-\n\n"

	for i, choice := range menuItems {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}

func main() {
	program := tea.NewProgram(initializeStartMenuModel())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unskyld! Something went wrong >w< here it is: %v", err)
		os.Exit(1)
	}
}
