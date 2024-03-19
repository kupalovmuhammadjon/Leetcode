package main

import (
	"fmt"
)

func main() {
	ls := []int{2, 1}
	fmt.Println(ls)
	ls = ls[:len(ls)-1]
	fmt.Println(ls)
	ls = ls[:len(ls)-2]
	fmt.Println(ls)

}
