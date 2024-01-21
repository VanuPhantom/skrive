package logic

import (
	"os"
	"time"
)

type Dose struct {
	Time      time.Time
	Quantity  string
	Substance string
	Route     string
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

func Load() ([]Dose, error) {
	if bytes, err := os.ReadFile("doses.dat"); err != nil {
		return nil, err
	} else {

		raw := string(bytes)

		return decode(raw)
	}
}
