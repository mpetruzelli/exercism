package hamming

import "errors"

func Distance(a, b string) (int, error) {
	shouldErr := len(a) - len(b)
	if shouldErr < 0 || shouldErr > 0 {
		return -1, errors.New("Lengths differ")
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
