func isPossibleToSplit(nums []int) bool {

	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}

	for _, k := range nums {
		if mp[k] > 2 {
			return false
		}
	}

	return true
}