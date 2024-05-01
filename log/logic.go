package log

import (
	"time"

	"skrive/data"

	tea "github.com/charmbracelet/bubbletea"
)

type logMsg struct {
	success bool
}

func timeFromOffset(offset int) time.Time {
	return time.Now().Add(time.Duration(-offset) * time.Minute)
}

func log(quantity string, substance string, route string, offset int) tea.Cmd {
	return func() tea.Msg {
		dose := data.Dose{
			Time:      timeFromOffset(offset),
			Quantity:  quantity,
			Substance: substance,
			Route:     route,
		}

		err := data.ApplicationStorage.Append(dose)

		return logMsg{
			success: err == nil,
		}
	}
}
