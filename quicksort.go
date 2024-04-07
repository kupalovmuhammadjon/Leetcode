package main

import (
	"fmt"
	"math/rand"
)

var count int = 0

func main() {
	// arr := []int{1, 2, 32, 3, 4, 345, 5, 76, 87, 8, 7, 65, 45, 45, 6, 54, 4, 6, 45, 4, 54, 545, 454, 54545, 454, 5454, 545}
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	// fmt.Println(quickSort(arr))
	shuffle(arr)
	fmt.Println(len(arr))
	fmt.Println(count)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	count++ // Increment count for entering the function
	pivot := medianOfThree(arr[0], arr[len(arr)/2], arr[len(arr)-1])
	less := []int{}
	greater := []int{}
	for i := 0; i < len(arr); i++ {
		count++ // Increment count for each iteration
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			greater = append(greater, arr[i])
		}
		// Elements equal to the pivot are ignored
	}
	res := append(quickSort(less), pivot)
	res = append(res, quickSort(greater)...)
	count += 2 // Increment count for the two append operations
	return res
}

// Function to find the median of three numbers
func medianOfThree(a, b, c int) int {
	if a < b {
		if b < c {
			return b
		} else if a < c {
			return c
		} else {
			return a
		}
	} else {
		if a < c {
			return a
		} else if b < c {
			return c
		} else {
			return b
		}
	}
}
func shuffle(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	quickSort(arr)
}

// func quickSort(arr []int) []int {
// 	count++
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	pivot := arr[len(arr)/2]
// 	less := []int{}
// 	greater := []int{}
// 	for i := 1; i < len(arr); i++ {
// 		count++
// 		if arr[i] <= pivot {
// 			less = append(less, arr[i])
// 		} else {
// 			greater = append(greater, arr[i])
// 		}
// 	}
// 	res := append(quickSort(less), pivot)
// 	res = append(res, quickSort(greater)...)
// 	count += 2
// 	return res
// }
