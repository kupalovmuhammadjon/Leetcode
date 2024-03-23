package main

import (
	"fmt"
)

func main() {
	ls := "123456"
	fmt.Println(ls[:len(ls)-1])
	ls = ls[1:2]

}
