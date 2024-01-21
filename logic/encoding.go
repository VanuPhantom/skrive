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
