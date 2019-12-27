package scrabble

import (
	"strings"
)

type pointStruct struct {
	letters string
	score   int
}

var scoring = []pointStruct{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

func Score(input string) int {
	total := 0
	for _, letter := range input {
		for _, pointList := range scoring {
			if strings.ContainsAny(pointList.letters, strings.ToUpper(string(letter))) {
				total = total + pointList.score
			}
		}

	}
	return total
}
