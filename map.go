package main

import "fmt"

func main(){
	s := "qwerty"
	t := "ytrewq"
	
	fmt.Println(isValidAnagram(s, t))
	
}
func isValidAnagram(s string, t string) bool{
	sfreq := map[rune]int{}
	tfreq := map[rune]int{}
	for _, char := range s{
		sfreq[char]++
	}
	matched := 0
	for _, char := range t{
		tfreq[char]++
		if tfreq[char] == sfreq[char]{
			matched++
		}
	}
	return len(sfreq) == matched
}

// nums := []int{1,2,3,4,5,6,7,8,9,7,5,4,3,2,2,1,4,5,7,8,9,9,7,5}
// 	freq := map[int]int{}
  
// 	mostFrequentNumber := [2]int{0, 0}
// 	for _, num := range nums{
// 	  freq[num]++
// 	  if mostFrequentNumber[1] < freq[num]{
// 		mostFrequentNumber[1] = freq[num]
// 		mostFrequentNumber[0] = num
// 	  }
// 	}
// 	fmt.Println(freq)
// 	fmt.Println(mostFrequentNumber)

// 	fmt.Println("String problem")
// 	sfreq := map[string]int{}
// 	s := "vfgbhjkiujyhtgfthyujytgrftHYUJytgrgthy"
// 	for i := 'a'; i <= 'z'; i++{
// 		sfreq[string(i)] = 0
// 	}
// 	for _, symb := range s{
// 		sfreq[string(symb)]++
// 	}
// 	fmt.Println(sfreq)