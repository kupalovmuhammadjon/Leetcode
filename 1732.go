func largestAltitude(gain []int) int {
	maxAltitude := 0
	curAltitude := 0
	for _, v := range gain {
		curAltitude += v
		if maxAltitude < curAltitude {
			maxAltitude = curAltitude
		}
	}
	return maxAltitude
}