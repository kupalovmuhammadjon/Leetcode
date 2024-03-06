func wordSubsets(words1 []string, words2 []string) []string {
	res := []string{}
	sub := counter(words2)
	for i := 0; i < len(words1); i++ {
		lamp := true
		for symb, count := range sub {
			if !(strings.Count(words1[i], string(symb)) >= count) {
				lamp = false
				break
			}
		}
		if lamp {
			res = append(res, words1[i])
		}
	}

	return res
}

func counter(words2 []string) map[rune]int {
	res := map[rune]int{}
	for j := 0; j < len(words2); j++ {
		sub := map[rune]int{}
		for _, c := range words2[j] {
			sub[c]++
		}
		for _, c := range words2[j] {
			if res[c] == 0 {
				res[c] = sub[c]
			} else if res[c] < sub[c] {
				res[c] = sub[c]
			}
		}
	}
	return res
}