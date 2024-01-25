package logic

import (
	"log"
	"os"
	"time"
)

type Dose struct {
	Time      time.Time
	Quantity  string
	Substance string
	Route     string
}

var Path = "/opt/homebrew/var/skrive/doses.dat"

func (d Dose) Log() error {
	file, err := os.OpenFile(Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Println(err.Error())
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
	if bytes, err := os.ReadFile(Path); err != nil {
		return nil, err
	} else {

		raw := string(bytes)

		return decode(raw)
	}
}
