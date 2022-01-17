package utils

import (
	"strings"
)

// MakeExcited transforms a sentence to all caps with an exclamation point
func MakeExcited(str string) string {
	return strings.ToUpper(str) + "!"
}
