package main

import "fmt"

func main() {
	fmt.Println(countDigits(10))
}

func countDigits(num int) int {
	cpy := num
	count := 0

	for cpy > 0 {
		if num  % (cpy%10) == 0 {
			count++
		}
		cpy /= 10
	}

	return count
}
