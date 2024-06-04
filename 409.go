func longestPalindrome(s string) int {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++{
        freq[s[i]]++
    }
    oddOne := 0
    length := 0
    for k := range freq{
        if freq[k] % 2 == 0{
            length += freq[k] 
        }else if freq[k] % 2 != 0 && freq[k] > 2{
            oddOne = 1
            length += freq[k] - 1
        }else{
            oddOne = 1
        }
    }

    return length + oddOne
}