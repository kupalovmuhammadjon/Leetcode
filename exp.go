package main

import (
	"fmt"
)

func main() {
	ls := "12"
	fmt.Println(ls)
	ls = ls[1:]
	fmt.Println(ls)
	ls = ls[1:]
	fmt.Println(ls)

}
