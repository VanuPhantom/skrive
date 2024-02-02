package view

import (
	"skrive/logic"

	tea "github.com/charmbracelet/bubbletea"
)

type successfulLoadMsg struct {
	doses []logic.Dose
}

type failedLoadMsg struct {
	err error
}

func load() tea.Msg {
	if doses, err := logic.Load(); err != nil {
		return failedLoadMsg{
			err,
		}
	} else {
		return successfulLoadMsg{
			doses,
		}
	}
}
