package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/akamensky/argparse"
)

var subcommand *string
var parser = argparse.NewParser("skrive", "Log doses via the terminal")
var fileArg = parser.String("f", "file", &argparse.Options{Required: false, Help: "Set dosage file path"})

func parse() error {
	subcommands := []string{}

	var remainingArgs []string

	if len(os.Args) >= 2 && slices.Contains(subcommands, os.Args[1]) {
		subcommand = &os.Args[1]
		remainingArgs = os.Args[2:]
		slices.Insert(remainingArgs, 0, "skrive")
	} else {
		remainingArgs = os.Args
	}

	return parser.Parse(remainingArgs)
}

func printHelp(err error) {
	fmt.Print(parser.Usage(err))
}
