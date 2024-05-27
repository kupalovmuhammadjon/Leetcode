func numsSameConsecDiff(n int, k int) []int {
    numbers := []int{}
    np := map[int]int{}
    var backtrack func(curNum, length int)
    backtrack = func(curNum, length int){
        if length == n{
            np[curNum]++
            numbers = append(numbers, curNum)
            return
        }

        if curNum % 10 + k < 10{
            curNum = curNum * 10 +curNum % 10 + k
            length++  
            backtrack(curNum, length)
            length--
            curNum /= 10
        }
        if curNum % 10 + k  != curNum % 10 - k && curNum % 10 - k >= 0{
            curNum = curNum * 10 + curNum % 10 - k
            length++
            backtrack(curNum, length)
            length--
            curNum /= 10
        }
        
    }
    for i := 1; i < 10; i++{
        backtrack(i, 1)
    }
    fmt.Println(np)
    return numbers
}
