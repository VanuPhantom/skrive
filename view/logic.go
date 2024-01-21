package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"skrive.vanu.dev/logic"
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
