package reverse

import (
	"strings"
)

func Reverse(input string) string {
	size := len(input)
	reverser := make([]string, size)
	if size == 0 {
		return input
	}
	for _, rune := range input {
		size--
		reverser[size] = string(rune)
	}
	var output string
	output = strings.Join(reverser, "")
	return output
}
