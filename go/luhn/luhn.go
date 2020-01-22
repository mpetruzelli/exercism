package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(input string) bool {
	input = strings.Join(strings.Split(input, " "), "")

	if len(input) <= 1 {
		return false
	}

	var inputInts []int
	for _, v := range input {
		converted, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		inputInts = append(inputInts, converted)
	}

	for index := len(inputInts) - 2; index >= 0; index -= 2 {
		doubled := inputInts[index] * 2
		if doubled > 9 {
			doubled = doubled - 9
		}
		inputInts[index] = doubled
	}
	var total int
	for _, v := range inputInts {
		total = total + v
	}
	fmt.Println(total)
	if total%10 == 0 {
		return true
	}
	return false
}
