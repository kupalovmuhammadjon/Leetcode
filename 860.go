func lemonadeChange(bills []int) bool {
	change := map[int]int{}
	for _, v := range bills {
		if v == 5 {
			change[v]++
			continue
		} else if v == 10 {
			change[5]--
			if change[5] < 0 {
				return false
			}
			change[10]++
		} else if v == 20 {
			if change[10] >= 1 {
				change[10] -= 1
				change[5] -= 1
				if change[5] < 0 {
					return false
				}
			} else if change[5] >= 3 {
				change[5] -= 3
			} else {
				return false
			}
		}

	}
	return true
}