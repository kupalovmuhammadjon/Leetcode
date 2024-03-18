func findDifference(nums1 []int, nums2 []int) [][]int {
	ans := [][]int{{}, {}}
	mp1 := map[int]int{}
	mp2 := map[int]int{}

	for _, v := range nums1 {
		mp1[v]++
	}
	for _, v := range nums2 {
		mp2[v]++
	}
	for k := range mp1 {
		if mp2[k] == 0 {
			ans[0] = append(ans[0], k)
		}
	}
	for k := range mp2 {
		if mp1[k] == 0 {
			ans[1] = append(ans[1], k)
		}
	}
	return ans
}