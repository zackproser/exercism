// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

const (
	SHOUT_RESPONSE            = "Whoa, chill out!"
	SHOUTED_QUESTION_RESPONSE = "Calm down, I know what I'm doing!"
	QUESTION_RESPONSE         = "Sure."
	STATEMENT_RESPONSE        = "Whatever."
	SILENCE_RESPONSE          = "Fine. Be that way!"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmed := strings.Trim(remark, " ")
	switch {
	case IsShouted(trimmed) && IsQuestion(trimmed):
		return SHOUTED_QUESTION_RESPONSE
	case IsShouted(trimmed):
		return SHOUT_RESPONSE
	case IsQuestion(trimmed):
		return QUESTION_RESPONSE
	case IsSilence(trimmed):
		return SILENCE_RESPONSE
	default:
		return STATEMENT_RESPONSE
	}
}

// IsShoutedQuestion returns true if the remark is BOTH shouted and a question
func IsShoutedQuestion(remark string) bool {
	return IsShouted(remark) && IsQuestion(remark)
}

// IsShouted returns true if every character in the provided string is capitalized
func IsShouted(remark string) bool {
	return (remark == strings.ToUpper(remark)) && (remark != strings.ToLower(remark))
}

// IsQuestion returns true if the provided string ends with a question mark
func IsQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

// IsSilence returns true if the remark is empty
func IsSilence(remark string) bool {
	return strings.ContainsAny(remark, "\t") || len(remark) == 0
}
