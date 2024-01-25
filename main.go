package main

import (
	"fmt"
	"log"
	"os"

	"skrive.vanu.dev/logic"
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

	if len(os.Args) >= 2 {
		logic.Path = os.Args[1]
	} else {
		err := os.MkdirAll("/opt/homebrew/var/skrive", os.ModePerm)
		if err != nil {
			log.Println(err.Error)
		}
	}

	program := tea.NewProgram(startMenu.InitializeModel())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unskyld! Something went wrong >w< here it is: %v", err)
		os.Exit(1)
	}
}
