package fs

import (
	"fmt"
	"skrive/data"
	"strconv"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func escape(raw string) string {
	escaped := strings.ReplaceAll(raw, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, ";", "\\;")

	return escaped
}

func encode(d data.Dose) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;", escape(string(d.Id)), escape(strconv.FormatInt(d.Time.Unix(), 10)), escape(d.Quantity), escape(d.Substance), escape(d.Route))
}

const (
	ENDS_WITH_ESCAPE = iota
	BAD_TIME
	UNKNOWN_HEADER
)

type DecodeError struct {
	Kind    int
	context *string
}

func (e DecodeError) Error() string {
	switch e.Kind {
	case ENDS_WITH_ESCAPE:
		return "The raw string ends with an escape character."
	case BAD_TIME:
		return "Failed to parse a time stamp."
	case UNKNOWN_HEADER:
		return fmt.Sprintf("Unknown header encountered when reading file: %s", *e.context)
	default:
		panic(fmt.Sprintf("DecodeError.Error has no case for error code %d.", e.Kind))
	}
}

type fieldParser func([]string) (*data.Dose, error)

func parse(raw string, fieldCount int, parseFields fieldParser) ([]data.Dose, error) {
	runes := []rune(raw)
	doses := make([]data.Dose, 0)

	for j := 0; j < len(runes); {
		sections := make([]string, fieldCount)
	sectionLoop:
		for i := 0; i < fieldCount; i++ {
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

		dose, err := parseFields(sections)

		if err != nil {
			return nil, err
		} else {
			doses = append(doses, *dose)
		}
	}

	return doses, nil
}

func parseHeaderless(raw string) ([]data.Dose, error) {
	return parse(raw, 4, func(sections []string) (*data.Dose, error) {
		unix, err := strconv.ParseInt(sections[0], 10, 64)

		if err != nil {
			return nil, DecodeError{Kind: BAD_TIME}
		}

		id, err := gonanoid.New()

		if err != nil {
			return nil, err
		}

		return &data.Dose{
			Id:        data.Id(id),
			Time:      time.Unix(unix, 0),
			Quantity:  sections[1],
			Substance: sections[2],
			Route:     sections[3],
		}, nil
	})
}

func parseVersion1(raw string) ([]data.Dose, error) {
	return parse(raw, 5, func(sections []string) (*data.Dose, error) {
		unix, err := strconv.ParseInt(sections[1], 10, 64)

		if err != nil {
			return nil, DecodeError{Kind: BAD_TIME}
		}

		return &data.Dose{
			Id:        data.Id(sections[0]),
			Time:      time.Unix(unix, 0),
			Quantity:  sections[2],
			Substance: sections[3],
			Route:     sections[4],
		}, nil
	})
}

func decode(raw string) ([]data.Dose, error) {
	if !strings.HasPrefix(raw, "Version:") {
		return parseHeaderless(raw)
	} else if strings.HasPrefix("Version:1\n", raw) {
		return parseVersion1(strings.TrimPrefix(raw, "Version:1\n"))
	} else {
		return nil, DecodeError{Kind: UNKNOWN_HEADER, context: &strings.SplitN(raw, "\n", 1)[0]}
	}
}
