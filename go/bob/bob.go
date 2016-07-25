// Lackadaisical teenager simulator.
package bob

import "unicode"
import "strings"

const testVersion = 2

// Send a message to Bob and get a response.
func Hey(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		// If nothing is said...
		return "Fine. Be that way!"
	}
	if isShouting(s) {
		return "Whoa, chill out!"
	}
	// Check ending punctuation.
	p := s[len(s)-1:]
	if p == "?" {
		// Question.
		return "Sure."
	} else {
		// Default response.
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
		if unicode.IsUpper(r) {
			// If the characters are capital, it's shouting.
			chill++
		} else if unicode.IsLower(r) {
			// Any lowercase letter means it's not shouting.
			return false
		}
	}
	if chill > 0 {
		// If we had at least one capital letter and no lowercase,
		// it's shouting.
		return true
	} else {
		return false
	}
}
