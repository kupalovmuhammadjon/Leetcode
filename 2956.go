func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := []int{}
	c1 := 0
	c2 := 0
	for i := 0; i < len(nums1); i++ {
		if contains(nums2, nums1[i]) {
			c1++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if contains(nums1, nums2[i]) {
			c2++
		}
	}
	res = append(res, c1)
	res = append(res, c2)
	return res
}

func contains(arr []int, t int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == t {
			return true
		}
	}
	return false
}