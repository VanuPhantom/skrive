package log

import (
	"fmt"
	"os"
	"skrive/logic"
	"time"
)

func Invoke(arguments []string) error {
	if len(arguments) != 3 {
		fmt.Println("Must either provide either 0 or 3 separate arguments (quantity, substance, & route)")
		os.Exit(1)
	}

	log(arguments[0], arguments[1], arguments[2], 0)

	dose := logic.Dose{
		Time:      time.Now(),
		Quantity:  arguments[0],
		Substance: arguments[1],
		Route:     arguments[2],
	}

	if err := dose.Log(); err != nil {
		return err
	}

	fmt.Printf("Logged %s of %s, taken via route %s\n", arguments[0], arguments[1], arguments[2])
	return nil
}
