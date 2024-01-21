package logic

import (
	"fmt"
	"os"
)

type Dose struct {
	Quantity  string
	Substance string
	Route     string
}

func (d Dose) encode() string {
	// TODO: Add escaping
	return fmt.Sprintf("%s;%s;%s;", d.Quantity, d.Substance, d.Route)
}

func (d Dose) Log() error {
	file, err := os.OpenFile("doses.dat", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	if _, err := file.WriteString(d.encode() + "\n"); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
