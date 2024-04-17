package main 

import "fmt"


func main(){
	bulbs := []int{0, 1, 0, 1, 1, 0, 1, 1}
	flipped := false

	for i := 0; i < len(bulbs); i++ {
		if bulbs[i] == 0{
			if flipped {
				bulbs[i] = 1
			}else{
				flipped = true
				bulbs[i] = 1
			}
		}else{
			if flipped{
				flipped = false
				bulbs[i] = 0
			}
		}
	}
	fmt.Println(bulbs)
}