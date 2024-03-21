func lengthOfLongestSubstring(s string) int {
	mx := 0
	start := 0
	mp := map[byte]int{}
	end := 0
	for end = 0; end < len(s); end++ {
		mp[s[end]]++
		for mp[s[end]] > 1 {
			key := s[start]
			mp[key]--
			start++
		}
		mx = max(mx, end-start+1)
	}

	return mx
}
