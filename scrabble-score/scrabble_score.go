package scrabble

import (
	"strings"
)

// LetterScores holds the mapping of characters to their values
type LetterScores struct {
	m map[string]int
}

// NewLetterScores returns a pointer to an instantiated LetterScores struct ready to use
func NewLetterScores() *LetterScores {
	l := &LetterScores{
		m: make(map[string]int),
	}
	l.buildMap()
	return l
}

// buildMap assigns the alphabet characters to their corresponding value in the map
func (l LetterScores) buildMap() {
	worth1 := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	worth2 := []string{"D", "G"}
	worth3 := []string{"B", "C", "M", "P"}
	worth4 := []string{"F", "H", "V", "W", "Y"}
	worth5 := []string{"K"}
	worth8 := []string{"J", "X"}
	worth10 := []string{"Q", "Z"}

	l.populateWithValue(worth1, 1)
	l.populateWithValue(worth2, 2)
	l.populateWithValue(worth3, 3)
	l.populateWithValue(worth4, 4)
	l.populateWithValue(worth5, 5)
	l.populateWithValue(worth8, 8)
	l.populateWithValue(worth10, 10)
}

// populdateWithValue accepts a slice of characters and a value to associate with each
// character in the underlying map
func (l LetterScores) populateWithValue(members []string, value int) {
	for _, e := range members {
		l.m[strings.ToLower(e)] = value
	}
}

// lookupValue takes a single character and returns its associated scrabble value
func (l LetterScores) lookupValue(c string) int {
	return l.m[strings.ToLower(c)]
}

// Score takes in the scrabble word and returns its total score value by adding
// the values for each of its member characters
func Score(word string) int {
	l := NewLetterScores()
	finalScore := 0
	for _, c := range word {
		val := l.lookupValue(string(c))
		finalScore += val
	}
	return finalScore
}
