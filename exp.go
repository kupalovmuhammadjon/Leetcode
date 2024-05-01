package main

import (
	"fmt"
)

func main() {
	s := "Salom"
	a := 3
	makeCub(&a)
	changeString(&s)
	fmt.Println(a)
	fmt.Println(s)
	
}
func makeCub(a *int){
	*a = *a * *a * *a
}

func changeString(s *string){
	*s = *s + *s + *s
}
