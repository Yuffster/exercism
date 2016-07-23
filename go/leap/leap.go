package leap

const testVersion = 2

func IsLeapYear(year int) bool {
	// Returns true for any year that's divisble by four.
	// Except any year divisible by 100.
	if year%100 == 0 {
		// Unless the year is also evenly divisible by 400.
		if year%400 == 0 {
			return true
		}
		return false
	} else {
		return (year%4 == 0)
	}
}