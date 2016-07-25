package secret

// Returns a secret handshake based on binary interpretation of the given
// int.
func Handshake(n int) []string {
	if n <= 0 {
		return []string{}
	}
	codes := []string{"wink", "double blink", "close your eyes", "jump"}
	out := []string{}
	for i, c := range codes {
		if getBit(n, i+1) {
			out = append(out, c)
		}
	}
	if getBit(n, 5) {
		// Reverse the order if the fifth bit if flipped.
		return reverse(out)
	}
	return out
}

// Returns true if the binary bit at the given place in an integer
// is 1.
func getBit(n, place int) bool {
	mask := uint(1) << uint(place-1)
	return uint(n) & mask > 0
}

// Algorithmic reversal.
func reverse(arr []string) []string {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}