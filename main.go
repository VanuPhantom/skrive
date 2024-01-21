package main

import (
	"fmt"
	"os"

	"skrive.vanu.dev/startMenu"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	program := tea.NewProgram(startMenu.InitializeModel())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unskyld! Something went wrong >w< here it is: %v", err)
		os.Exit(1)
	}
}
