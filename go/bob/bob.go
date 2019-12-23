/*
Package bob prints simple text based on user input, because AI is just "Add Ifs" right?

User inputs that will elicit specific output:

Simple question: 'Sure.'
Upper-case statement: 'Whoa, chill out!'
Upper-case question: 'Calm down, I know what I'm doing!'
No input: 'Fine. Be that way!'
All other input: 'Whatever.'
*/
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark) && !isYelling(remark):
		return "Sure."
	case !isQuestion(remark) && isYelling(remark):
		return "Whoa, chill out!"
	case len(remark) == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}

func isQuestion(input string) bool {
	if strings.HasSuffix(input, "?") {
		return true
	}
	return false
}

func isYelling(input string) bool {
	if strings.ToUpper(input) == input && input != strings.ToLower(input) {
		return true
	}
	return false
}
