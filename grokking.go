package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 9; i < 129; i++ {
		arr = append(arr, i)
	}
	// fmt.Println(sum(arr))
	// fmt.Println(count(arr))
	// fmt.Println(max(arr, 0))
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 1, 0))
	fmt.Println()
	// fmt.Println(binarySearch(arr, 0, len(arr)-1, 5, 0))
	// fmt.Println(binarySearch(arr, 0, len(arr)-1, 6, 0))
	// fmt.Println(binarySearch(arr, 0, len(arr)-1, 7, 0))
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 64, 0))
	fmt.Println()
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 128, 0))
}
func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[len(arr)-1] + sum(arr[:len(arr)-1])
}

func count(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	return 1 + count(arr[:len(arr)-1])
}

func max(arr []int, mx int) int {
	if len(arr) == 1 {
		if arr[0] > mx {
			return mx
		}
		return mx
	}
	if mx < arr[len(arr)-1] {
		mx = arr[len(arr)-1]
	}
	return max(arr[:len(arr)-1], mx)
}

func binarySearch(arr []int, l, r, tar, s int) int {
	if l > r || len(arr) < r {
		return -1
	}
	mid := (l + r) / 2
	fmt.Println(arr[mid])
	guess := arr[mid]
	if guess == tar {
		return s
	} else if guess < tar {
		l = mid + 1
	} else {
		r = mid - 1
	}
	return binarySearch(arr, l, r, tar, s+1)
}
