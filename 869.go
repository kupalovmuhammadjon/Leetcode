func reorderedPowerOf2(n int) bool {
	mp := map[int]int{}
	d := n
	st := strconv.Itoa(n)
	for _, v := range st {
		mp[int(v-'0')]++
	}
	for i := 1; i <= d*15; i *= 2 {
		if isMatch(mp, i) {
			return true
		}
	}

	return false
}

func isMatch(mp map[int]int, a int) bool {
	ap := map[int]int{}

	for a > 0 {
		ap[a%10]++
		a /= 10
	}

	for k := range ap {
		if mp[k] != ap[k] {
			return false
		}
	}

	for k := range mp {
		if mp[k] != ap[k] {
			return false
		}
	}

	return true
}
