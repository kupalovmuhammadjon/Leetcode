package main

import "fmt"

func main() {
	ls := "12"
	ls = ls[:1] + ls[2:]
	fmt.Println(ls)

}
