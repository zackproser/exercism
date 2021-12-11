package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	noLettersRepeated := true

	seen := make(map[string]bool)

	for _, c := range word {
		key := strings.ToLower(string(c))
		if _, ok := seen[key]; ok {
			noLettersRepeated = false
			break
		}
		// Only store characters as seen if they are letters (skipping hyphens and spaces)
		if unicode.IsLetter(c) {
			seen[key] = true
		}
	}
	return noLettersRepeated
}
