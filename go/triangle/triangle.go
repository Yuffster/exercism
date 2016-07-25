// Utility methods concerning triangles.
package triangle

import "math"

const testVersion = 2

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// Returns a Kind type depending on what kind of triangle we have.
//
// Refer to the constant list (NaT, Equ, Iso, Sca) for information
// regarding return values.
func KindFromSides(a, b, c float64) Kind {
	if a == 0 && b == 0 && c == 0 {
		// No angles, no triangle.
		return NaT
	}
	if notNumbers(a, b, c) {
		// Triangles need numbers.
		return NaT
	}
	if a+b < c || c+a < b || b+c < a {
		// The sum of the lengths of any two sides of a triangle always
		// exceeds or is equal to the length of the third side, a principle
		// known as the triangle inequality.
		return NaT
	}
	if a == b && a == c {
		// An equilateral triangle has three equal sides.
		return Equ
	}
	if a == b || b == c || c == a {
		// An isosceles triangle has two sides of equal length.
		return Iso
	}
	if a != b && b != c && c != a {
		// A scalene triangle has three unequal sides.
		return Sca
	}
	return NaT
}

func notNumbers(ns ...float64) bool {
	for _, v := range ns {
		if math.IsInf(v, -1) || math.IsInf(v, 1) || math.IsNaN(v) {
			return true
		}
	}
	return false
}
