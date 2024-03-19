func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	mp1 := map[rune]int{}
	mp2 := map[rune]int{}
	set1 := []string{}
	set2 := []string{}
	freq1 := []int{}
	freq2 := []int{}

	for _, v := range word1 {
		mp1[v]++
	}
	for _, v := range word2 {
		mp2[v]++
	}
	lamp := true
	for k := range mp1 {
		set1 = append(set1, string(k))
		freq1 = append(freq1, mp1[k])
		if mp1[k] != mp2[k] {
			lamp = false
		}
	}
	for k := range mp2 {
		set2 = append(set2, string(k))
		freq2 = append(freq2, mp2[k])
		if mp1[k] != mp2[k] {
			lamp = false
		}
	}
	if lamp {
		return true
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	sort.Strings(set1)
	sort.Strings(set2)
	for i := 0; i < len(set1); i++ {
		if set1[i] != set2[i] {
			return false
		}
	}

	return true
}