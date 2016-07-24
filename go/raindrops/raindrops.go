// Raindrop translation library.
package raindrops

import "fmt"

const testVersion = 2

// Converts an int to a pattern of raindrops.
func Convert(i int) string {
	s := ""
	// If the number contains 3 as a factor, output 'Pling'.
	if i%3 == 0 {
		s += "Pling"
	}
	// If the number contains 5 as a factor, output 'Plang'.
	if i%5 == 0 {
		s += "Plang"
	}
	// If the number contains 7 as a factor, output 'Plong'.
	if i%7 == 0 {
		s += "Plong"
	}
	// If the number does not contain 3, 5, or 7 as a factor,
	// just pass the number's digits straight through.
	if i%7 > 0 && i%3 > 0 && i%5 > 0 {
		return fmt.Sprintf("%d", i)
	}
	return s
}
