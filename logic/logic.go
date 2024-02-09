package logic

import (
	"log"
	"os"
	"path"
	"sort"
	"time"
)

type Dose struct {
	Time      time.Time
	Quantity  string
	Substance string
	Route     string
}

var dosageFilePath string

func Setup() error {
	if len(os.Args) > 2 {
		dosageFilePath = os.Args[1]
		return nil
	}

	if value, isDefined := os.LookupEnv("SKRIVE_DOSES_PATH"); isDefined {
		dosageFilePath = value
		return nil
	}

	if dirname, err := os.UserHomeDir(); err != nil {
		return homePathError{}
	} else {
		dosageFilePath = path.Join(dirname, "doses.dat")
		return nil
	}
}

type homePathError struct{}

func (e homePathError) Error() string {
	return "Could not find home directory"
}

func (d Dose) Log() error {
	file, err := os.OpenFile(dosageFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

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
	if bytes, err := os.ReadFile(dosageFilePath); err != nil {
		return nil, err
	} else {

		raw := string(bytes)

		if doses, err := decode(raw); err != nil {
			return nil, err
		} else {
			sort.Slice(doses, func(i, j int) bool {
				return doses[i].Time.Unix() > doses[j].Time.Unix()
			})

			return doses, nil
		}
	}
}
