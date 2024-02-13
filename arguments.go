package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/akamensky/argparse"
)

var subcommand *string
var parser = argparse.NewParser("skrive", "Log doses via the terminal")
var fileArg = parser.String("f", "file", &argparse.Options{Required: false, Help: "Set the doses file path"})
var helpFlag = parser.Flag("h", "help", &argparse.Options{Help: "Print help information"})
var positionalArguments = []string{}

func parse() error {
	// Our custom help flag is used for properly displaying usage information
	// The default behavior does not account for the usage of the subcommands
	//  defined in printHelp()
	parser.DisableHelp()
	subcommands := []string{"log"}

	var remainingArgs []string

	if len(os.Args) >= 2 && slices.Contains(subcommands, os.Args[1]) {
		subcommand = &os.Args[1]
		remainingArgs = slices.Insert(os.Args[2:], 0, "skrive")
	} else {
		remainingArgs = os.Args
	}

	positionalParameters := [4]*string{}

	if subcommand != nil && *subcommand == "log" {
		options := argparse.Options{Help: argparse.DisableDescription}

		for i := range positionalParameters {
			positionalParameters[i] = parser.StringPositional(&options)
		}
	}

	var err = parser.Parse(remainingArgs)

	for i := range positionalParameters {
		if positionalParameters[i] == nil || *positionalParameters[i] == "" {
			continue
		}
		positionalArguments = slices.Insert(positionalArguments, i, *positionalParameters[i])
	}

	return err
}

func printHelp(err error) {
	parser.NewCommand("<blank>", "Open the TUI")
	parser.NewCommand("log", "Log a dose")
	fmt.Print(parser.Usage(err))
}
