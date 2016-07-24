package gigasecond
import "time"
import "math"

const testVersion = 4

// Returns one gigasecond (10^9 seconds) from the passed in time,
// so you can celebrate your birthgigasecond.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := math.Pow(10, 9)
	return t.Add(time.Duration(gigasecond) * time.Second)
}