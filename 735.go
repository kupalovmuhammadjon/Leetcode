func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, a := range asteroids {
		if a > 0 {
			stack = append(stack, a)
			continue
		}
		lamp := true
		for a < 0 && len(stack) != 0 && stack[len(stack)-1] > 0 {
			side := (stack[len(stack)-1] + a)
			if side > 0 {
				lamp = false
				break
			} else if side < 0 {
				stack = stack[:len(stack)-1]
			} else {
				lamp = false
				stack = stack[:len(stack)-1]
				break
			}
		}
		if lamp {
			stack = append(stack, a)
		}

	}
	return stack
}
