package main

import (
	"fmt"
	"os"

	"skrive/logic"
	"skrive/startMenu"

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

	err := logic.Setup()

	if err == nil {
		_, err = tea.
			NewProgram(startMenu.InitializeModel()).
			Run()
	}

	if err != nil {
		fmt.Println("Undskyld! Something went wrong >w< here it is: %v", err)
		os.Exit(1)
	}
}
