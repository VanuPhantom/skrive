package log

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"skrive.vanu.dev/logic"
)

type logMsg struct {
	success bool
}

func log(quantity string, substance string, route string) tea.Cmd {
	return func() tea.Msg {
		dose := logic.Dose{
			Time:      time.Now(),
			Quantity:  quantity,
			Substance: substance,
			Route:     route,
		}

		err := dose.Log()

		return logMsg{
			success: err == nil,
		}
	}
}
