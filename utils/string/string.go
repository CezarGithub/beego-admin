package string

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func Alphanumeric(str string) string {
	s := nonAlphanumericRegex.ReplaceAllString(str, "")
	return strings.ToLower(s)
}
