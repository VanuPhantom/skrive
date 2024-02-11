package main

import (
	"fmt"
	"os"

	"skrive/log"
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

	parseErr := parse()
	if parseErr != nil {
		printHelp(parseErr)
		os.Exit(1)
	} else if *helpFlag {
		printHelp(nil)
		os.Exit(0)
	}

	err := logic.Setup(*fileArg)

	if err == nil && subcommand != nil {
		handleSubcommands()
	}

	if err == nil {
		var model tea.Model
		if subcommand != nil && *subcommand == "log" {
			model, _ = log.InitializeModel(func() (tea.Model, tea.Cmd) {
				return model, tea.Quit
			})
		} else {
			model = startMenu.InitializeModel()
		}

		_, err = tea.
			NewProgram(model).
			Run()
	}

	handleIfError(err)
}

func handleSubcommands() {
	switch *subcommand {
	case "log":
		if len(positionalArguments) == 0 {
			// Handled in Bubbletea initialization code
			return
		}
		handleIfError(log.Invoke(positionalArguments))
	}
	os.Exit(0)
}

func handleIfError(err error) {
	if err == nil {
		return
	}
	fmt.Println("Undskyld! Something went wrong >w< here it is: %v", err)
	os.Exit(1)
}
