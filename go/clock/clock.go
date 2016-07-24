// Provides a struct (Clock) which keeps track of a signed integer of
// minutes.
package clock

import "fmt"
import "math"

const testVersion = 4

// In case our system of keeping time ever changes.
const hours = 24
const minutes = 60
const clockLimit = hours * minutes

// A clock is simply a struct that stores minutes as a base number,
// then computes the hour and minute values in a 24-hour clock when
// converted to a string.
//
// Internally, minutes can be positive or negative, and can include
// far longer than the single twenty-four hour cycle present in the
// final clock value.
type Clock struct {
	minutes int
}

// Creates a new clock based on the hour and minute specified.
//
// Hour and minute can be beyond the range of supported clock values
// and will be normalized.
func New(hour, minute int) Clock {
	return Clock{0}.Add(hour*minutes + minute)
}

// Converts the minutes stored in the clock to a string of
// HH:MM on a 24-hour clock, zero padded.
func (c Clock) String() string {
	m := c.minutes
	h := 0
	h = int(math.Floor(float64(m / minutes)))
	if h == 24 {
		h = 0
	}
	m %= minutes
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Adds (or subtracts) the designated number of minutes.
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	// Normalize minutes outside of standard clock range.
	if c.minutes < 0 {
		if c.minutes < clockLimit*-1 {
			c.minutes = clockLimit + (c.minutes % clockLimit)
		} else {
			c.minutes = clockLimit + c.minutes
		}
	} else if c.minutes > clockLimit {
		c.minutes %= clockLimit
	}
	return c
}