package fs

import (
	"fmt"
	"skrive/data"
	"strconv"
	"strings"
	"time"
)

func escape(raw string) string {
	escaped := strings.ReplaceAll(raw, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, ";", "\\;")

	return escaped
}

func encode(d data.Dose) string {
	return fmt.Sprintf("%s;%s;%s;%s;", escape(strconv.FormatInt(d.Time.Unix(), 10)), escape(d.Quantity), escape(d.Substance), escape(d.Route))
}

const (
	ENDS_WITH_ESCAPE = iota
	BAD_TIME
)

type DecodeError struct {
	Kind int
}

func (e DecodeError) Error() string {
	switch e.Kind {
	case ENDS_WITH_ESCAPE:
		return "The raw string ends with an escape character."
	case BAD_TIME:
		return "Failed to parse a time stamp."
	default:
		panic(fmt.Sprintf("DecodeError.Error has no case for error code %d.", e.Kind))
	}
}

func decode(raw string) ([]data.Dose, error) {
	runes := []rune(raw)
	doses := make([]data.Dose, 0)

	for j := 0; j < len(runes); {
		sections := [4]string{"", "", "", ""}
	sectionLoop:
		for i := 0; i < 4; i++ {
			for ; j < len(runes); j++ {
				character := runes[j]

				if character == '\\' {
					j++
					if j >= len(runes) {
						return nil, DecodeError{Kind: ENDS_WITH_ESCAPE}
					}
				} else if character == ';' {
					j++
					continue sectionLoop
				}

				sections[i] += string(runes[j])
			}
		}

		if runes[j] == '\n' {
			j++
		}

		unix, err := strconv.ParseInt(sections[0], 10, 64)

		if err != nil {
			return nil, DecodeError{Kind: BAD_TIME}
		}

		doses = append(doses,
			data.Dose{
				Time:      time.Unix(unix, 0),
				Quantity:  sections[1],
				Substance: sections[2],
				Route:     sections[3],
			})
	}

	return doses, nil
}
