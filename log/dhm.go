package log

import (
	"strconv"
	"strings"
)

func increment(raw_before_delimiter *string, field *int) error {
	val, err := strconv.Atoi(*raw_before_delimiter)
	*raw_before_delimiter = ""

	if err == nil {
		*field += val
	}

	return err
}

type emptyTimeError struct{}

func (e *emptyTimeError) Error() string {
	return "The time cannot be empty."
}

func parseTime(raw string) (int, error) {
	raw = strings.Trim(raw, " ")

	if len(raw) == 0 {
		return 0, &emptyTimeError{}
	}

	days := 0
	hours := 0
	minutes := 0

	raw_before_delimiter := ""

	for i := range raw {
		r := raw[i]
		var err error

		switch r {
		case 'd':
			err = increment(&raw_before_delimiter, &days)
		case 'h':
			err = increment(&raw_before_delimiter, &hours)
		case 'm':
			err = increment(&raw_before_delimiter, &minutes)
		default:
			raw_before_delimiter += string(r)
		}

		if err != nil {
			return 0, err
		}
	}

	if len(raw_before_delimiter) > 0 {
		err := increment(&raw_before_delimiter, &minutes)

		if err != nil {
			return 0, err
		}
	}

	hours += days * 24
	minutes += hours * 60

	return minutes, nil
}
