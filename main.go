package main

import (
	"fmt"
	"os"

	"skrive.vanu.dev/startMenu"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(startMenu.InitializeModel())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unskyld! Something went wrong >w< here it is: %v", err)
		os.Exit(1)
	}
}
