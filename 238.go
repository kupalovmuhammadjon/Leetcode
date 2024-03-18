func productExceptSelf(nums []int) []int {
	res := []int{}
	sm := 1
	zeros := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros++
			continue
		}
		sm *= nums[i]
	}
	if zeros == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				res = append(res, sm)
			} else {
				res = append(res, 0)
			}
		}
	} else if zeros > 1 {
		for i := 0; i < len(nums); i++ {
			res = append(res, 0)
		}
	} else {
		for i := 0; i < len(nums); i++ {
			res = append(res, sm/nums[i])
		}
	}
	return res
}