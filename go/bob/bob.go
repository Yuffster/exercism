// Lackadaisical teenager simulator.
package bob

import "unicode"
import "strings"

const testVersion = 2

// Send a message to Bob and get a response.
func Hey(s string) string {
	s = strings.TrimSpace(s)
	// If nothing is said...
	if len(s) == 0 {
		return "Fine. Be that way!"
	}
	if isShouting(s) {
		return "Whoa, chill out!"
	}
	// Check ending punctuation.
	p := s[len(s)-1:]
	// Statement.
	if p == "." {
		return "Whatever."
		// Question.
	} else if p == "?" {
		return "Sure."
		// Default response.
	} else {
		return "Whatever."
	}
}

// Check to see if the speaker is shouting.
//
// This is generally a good indication that the speaker needs to
// chill out.
func isShouting(s string) bool {
	chill := 0
	for _, c := range s {
		r := rune(c)
		// If the characters are capital, it's shouting.
		if unicode.IsUpper(r) {
			chill++
		// Any lowercase letter means it's not shouting.
		} else if unicode.IsLower(r) {
			return false
		}
	}
	// If we had at least one capital letter and no lowercase,
	// it's shouting.
	if chill > 0 {
		return true
	} else {
		return false
	}
}
