package main

import (
	"fmt"
)

func main() {
	// birthYear := 2005
	// currentYear := 2024
	// fmt.Println(currentYear - birthYear)
	// var b byte = 'с'
	var r rune = 'й'
	fmt.Println(r)
	// s := "Helo"
	// for i, _ := range s{
	// 	fmt.Println(s[i])
	// }
	// var s string = "Hello"
	// sInBytes := []byte(s)
	// sInRunes := []rune(s)
	// fmt.Println(string(sInBytes))
	// fmt.Println(string(sInRunes))

	var s string = "你好世界"
sInBytes := []byte(s)
sInRunes := []rune(s)
fmt.Println(sInBytes)
fmt.Println(string(sInBytes))
fmt.Println(string(sInRunes))

}
