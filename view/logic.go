package view

import (
	"skrive/data"

	tea "github.com/charmbracelet/bubbletea"
)

type successfulLoadMsg struct {
	doses []data.Dose
}

type failedLoadMsg struct {
	err error
}

func load() tea.Msg {
	if doses, err := data.ApplicationStorage.FetchAll(); err != nil {
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
	doses []data.Dose
}

func remove(id data.Id, doses []data.Dose) tea.Cmd {
	return func() tea.Msg {
		newDoses := make([]data.Dose, 0)

		for _, dose := range doses {
			if dose.Id != id {
				newDoses = append(newDoses, dose)
			}
		}

		if data.ApplicationStorage.DeleteDose(id) == nil {
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
