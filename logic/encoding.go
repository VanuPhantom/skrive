package logic

import (
	"fmt"
	"strings"
)

func escape(raw string) string {
	escaped := strings.ReplaceAll(raw, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, ";", "\\;")

	return escaped
}

func (d Dose) encode() string {
	return fmt.Sprintf("%s;%s;%s;", escape(d.Quantity), escape(d.Substance), escape(d.Route))
}

const (
	ENDS_WITH_ESCAPE = iota
)

type DecodeError struct {
	Kind int
}

func (e DecodeError) Error() string {
	switch e.Kind {
	case ENDS_WITH_ESCAPE:
		return "The raw string ends with an escape character."
	default:
		panic(fmt.Sprintf("DecodeError.Error has no case for error code %d.", e.Kind))
	}
}

func decode(raw string) ([]Dose, error) {
	runes := []rune(raw)
	doses := make([]Dose, 0)

	for j := 0; j < len(runes); {
		sections := [3]string{"", "", ""}
	sectionLoop:
		for i := 0; i < 3; i++ {
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

		doses = append(doses,
			Dose{
				Quantity:  sections[0],
				Substance: sections[1],
				Route:     sections[2],
			})
	}

	return doses, nil
}
