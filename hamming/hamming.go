package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	aChars := strings.Split(a, "")
	bChars := strings.Split(b, "")
	distance := 0

	if len(aChars) != len(bChars) {
		return distance, errors.New("Input strings must be of equal length")
	}

	for i, c := range aChars {
		if c != bChars[i] {
			distance++
		}
	}
	return distance, nil
}
