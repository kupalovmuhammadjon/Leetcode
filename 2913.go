func sumCounts(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)+1; j++ {
			count := setlen(nums[i:j])
			res += count * count
		}
	}
	return res
}

func setlen(temp []int) int {
	tp := []int{}
	for k := 0; k < len(temp); k++ {
		lamp := 1
		for h := 0; h < len(tp); h++ {
			if tp[h] == temp[k] {
				lamp = 0
			}
		}
		if lamp == 1 {
			tp = append(tp, temp[k])
		}

	}
	return len(tp)
}