package raindrops

import (
	"math"
	"strconv"
)

func Convert(input int) (output string) {
	var base = float64(input)
	if math.Mod(base, 3) == 0 {
		output = output + "Pling"
	}

	if math.Mod(base, 5) == 0 {
		output = output + "Plang"
	}

	if math.Mod(base, 7) == 0 {
		output = output + "Plong"
	}

	if len(output) == 0 {
		output = strconv.Itoa(input)
	}
	return output
}
