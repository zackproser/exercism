package pangram

import (
	"fmt"
	"strings"
)

// IsPangram returns true if the supplied string uses all letters in the alphabet at least once, and false if it does not
func IsPangram(input string) bool {
	if input == "" {
		return false
	}

	letterMap := make(map[string]int)

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// Build initial map, where each letter will have a count of 0
	for _, c := range alphabet {
		letterMap[c] = 0
	}

	for _, c := range input {
		letterMap[strings.ToLower(string(c))]++
	}

	allLettersSeenAtLeastOnce := true
	for k, v := range letterMap {
		if v < 1 {
			fmt.Printf("Letter %s seen only: %d", k, v)
			allLettersSeenAtLeastOnce = false
		}
	}
	return allLettersSeenAtLeastOnce
}
