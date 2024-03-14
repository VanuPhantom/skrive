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

type removeMsg struct {
	doses []logic.Dose
}

func remove(doses []logic.Dose, index int) tea.Cmd {
	return func() tea.Msg {
		newDoses := make([]logic.Dose, 0)
		newDoses = append(newDoses, doses[:index]...)
		newDoses = append(newDoses, doses[index+1:]...)

		if logic.Overwrite(newDoses) == nil {
			return removeMsg{
				doses: newDoses,
			}
		} else {
			return removeMsg{
				doses,
			}
		}
	}
}
