package log

import (
	"fmt"
	"os"
	"skrive/logic"
	"strconv"
	"time"
)

func Invoke(arguments []string) error {
	if len(arguments) < 3 && len(arguments) > 4 {
		fmt.Println("Must either provide either 0, 3 or 4 separate arguments (quantity, substance, route and - optionally - minutes since dose)")
		os.Exit(1)
	}

	time := time.Now()
	offsetDescription := ""

	if len(arguments) > 3 {
		value, err := strconv.Atoi(arguments[3])

		if err != nil {
			fmt.Println("Minutes since dose must be an integer!")
			os.Exit(1)
		}

		time = timeFromOffset(value)
		offsetDescription = fmt.Sprintf(" %d minutes ago", value)
	}

	dose := logic.Dose{
		Time:      time,
		Quantity:  arguments[0],
		Substance: arguments[1],
		Route:     arguments[2],
	}

	if err := dose.Log(); err != nil {
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
