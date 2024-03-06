func minimumLength(s string) int {
	res := s
	for i := 0; i < len(s); i++ {
		p := res[0]
		pind := 0
		sind := -1
		for j := 0; j <= len(res)/2; j++ {
			if p == res[j] {
				pind = j
			} else {
				break
			}
		}
		for j := len(res) - 1; j >= len(res)/2; j-- {
			if p == res[j] {
				sind = j
				// if  sind == pind{
				//     return len(res)
				// }
			} else {
				break
			}
			fmt.Println(res)
		}
		if sind == -1 {
			return len(res)
		}
		if len(res) == 2 && sind != -1 {
			return 0
		}
		res = res[pind+1 : sind]
		if len(res) == 0 {
			return 0
		}
	}

	return len(res)
}

func minimumLength(s string) int {
	left := 0
	right := len(s) - 1

	for left < right && s[left] == s[right] {
		current := s[left]
		for left <= right && s[left] == current {
			left++
		}
		// fmt.Println(left)
		// fmt.Println(right)
		for right >= left && s[right] == current {
			right--
			// fmt.Println(s[right])
		}
	}

	return right - left + 1
}
