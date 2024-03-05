func maxDistToClosest(seats []int) int {
	max := 0
	g := 0
	prv := false
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if !prv {
				if max < g {
					max = g
				}
			} else {
				if g%2 == 0 {
					if max < g/2 {
						max = g / 2
					}
				} else {
					if max < (g+1)/2 {
						max = (g + 1) / 2
					}
				}
			}
			g = 0
			prv = true
		} else {
			g++
		}
	}
	if max < g {
		max = g
	}
	return max
}