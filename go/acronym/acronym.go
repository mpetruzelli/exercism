// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9']+")
	processedString := reg.ReplaceAllString(s, " ")
	wordList := strings.Split(processedString, " ")
	firstLetter := make([]string, len(wordList))
	for i, word := range wordList {
		firstLetter[i] = string(word[0])
	}
	acronym := strings.ToUpper(strings.Join(firstLetter, ""))
	return acronym
}
