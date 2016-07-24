// Provides a Hamming distance utility.
package hamming

import "errors"

const testVersion = 4

// Returns the Hamming distance between two strings of equal length.
//
// This is calculated by the number of characters which differ between
// the two strings, dependent on their location in the sequence.
//
// If the strings are of an unequal length, an error will be returned.
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("String length must be equal.")
	}
	for i, _ := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
