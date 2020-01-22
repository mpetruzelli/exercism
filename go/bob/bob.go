/*
Package bob implements a simple chatbot.

bob responds to text input in the following ways:

Lower-cased question: "Sure."
Upper-cased question: "Calm down, I know what I'm doing!"
Upper-cased statement: "Whoa, chill out!"
Null text input: "Fine. Be that way!"
Other: "Whatever."

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := UserInput()
	fmt.Println(input)

}

// Responds to the user input as bob
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return ""
}

// Take user input and return that input
func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What's up?")
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	sanitizedInput := strings.Replace(input, "\n", "", -1)
	return sanitizedInput
}
