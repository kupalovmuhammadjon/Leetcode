func hasAlternatingBits(n int) bool {
	prev := -1

	for n > 0 {
		if prev == n%2 {
			return false
		} else {
			prev = n % 2
			n /= 2
		}
	}
	return true
}