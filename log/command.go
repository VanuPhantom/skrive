package log

import (
	"fmt"
	"os"
	"skrive/data"
	"time"
)

func Invoke(storage data.Storage, arguments []string) error {
	if len(arguments) < 3 && len(arguments) > 4 {
		fmt.Println("Usage: " +
			"skrive log [-f path to doses.dat] " +
			"[<quantity> <substance> <route> [time-spec]]")
		os.Exit(1)
	}

	time := time.Now()
	offsetDescription := ""

	if len(arguments) > 3 {
		value, err := parseTime(arguments[3])

		if err != nil {
			fmt.Println("Time argument must be like 1d2h3m or 123 (for 123 minutes)")
			os.Exit(1)
		}

		time = timeFromOffset(value)
		offsetDescription = fmt.Sprintf(" %d minutes ago", value)
	}

	dose := data.Dose{
		Time:      time,
		Quantity:  arguments[0],
		Substance: arguments[1],
		Route:     arguments[2],
	}

	if err := storage.Append(dose); err != nil {
		return err
	}

	fmt.Printf(
		"Logged %s of %s, taken via route %s%s\n",
		arguments[0],
		arguments[1],
		arguments[2],
		offsetDescription,
	)
	return nil
}
