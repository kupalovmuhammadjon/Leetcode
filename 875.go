func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	k := 0
	l := 1
	r := piles[len(piles)-1]

	for l <= r {
		m := (r + l) / 2
		if canKokoEat(piles, m, h) {
			k = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return k
}
func canKokoEat(piles []int, k, h int) bool {
	totalHour := 0
	for _, v := range piles {
		if v%k == 0 {
			totalHour += v / k
		} else {
			totalHour += v/k + 1
		}
	}
	if totalHour <= h {
		return true
	}
	return false
}