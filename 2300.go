func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}
	for _, v := range spells {
		if int64(v) >= success || success <= int64(v)*int64(potions[0]) {
			res = append(res, len(potions))
		} else if success%int64(v) == 0 {
			r := binarySearch(potions, success/int64(v))
			if r == -1 {
				res = append(res, 0)
			} else {
				res = append(res, len(potions)-r)
			}
		} else {
			r := binarySearch(potions, success/int64(v)+1)
			if r == -1 {
				res = append(res, 0)
			} else {
				res = append(res, len(potions)-r)
			}
		}
	}
	return res
}
func binarySearch(arr []int, tar int64) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if int64(arr[mid]) >= tar {
			if mid == 0 || int64(arr[mid-1]) < tar {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

