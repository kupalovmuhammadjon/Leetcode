package main

import "fmt"

var count int = 0

func main() {
	arr := []int{1, 2, 32, 3, 4, 345, 5, 76, 87, 8, 7, 65, 45, 45, 6, 54, 4, 6, 45, 4, 54, 545, 454, 54545, 454, 5454, 545}
	fmt.Println(quicksort(arr, 0, len(arr)-1))
	fmt.Println(len(arr))
	fmt.Println(count)
}
func quicksort(arr []int, start, finish int) []int {
	count++
	if start < finish {
		pivot := arr[(start+finish)/2]
		i := start
		j := finish
		for i <= j {
			count++
			for arr[i] < pivot {
				count++
				i++
			}
			for arr[j] > pivot {
				count++
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		quicksort(arr, start, j)
		quicksort(arr, i, finish)
	}
	return arr
}

// func quicksort(arr []int, start, finish int) []int {
// 	count++
// 	i := start
// 	j := finish
// 	middle := (start + finish) / 2
// 	for i <= j {
// 		count++
// 		for arr[i] < arr[middle] {
// 			count++
// 			i++
// 		}
// 		for arr[middle] < arr[j] {
// 			count++
// 			j--
// 		}
// 		if i <= j {

// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 			j--
// 		}
// 		if start < j {
// 			quicksort(arr, start, j)
// 		}
// 		if finish > i {
// 			quicksort(arr, i, finish)
// 		}
// 	}
// 	return arr
// }
