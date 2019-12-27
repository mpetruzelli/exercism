package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {
	input = strings.Replace(strings.ToLower(input), "-", "", -1)
	input = strings.Replace(input, " ", "", -1)
	for _, rune := range input {
		letter := strings.ToLower(string(rune))
		if strings.Count(input, letter) > 1 {
			return false
		}
	}
	return true
}
