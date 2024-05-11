func combinationSum3(k int, n int) [][]int {
    combinations := [][]int{}
    var backtrack func(combo []int, currentSum int, ind int)
    backtrack = func(combo []int, currentSum int, ind int){
        if len(combo) == k && currentSum == n{
            copyCombo := make([]int, k)
            copy(copyCombo, combo)
            combinations = append(combinations, copyCombo)
            return
        }else if len(combo) > k || currentSum > n{
            return
        }
        fmt.Println(ind)
        for i := ind; i < 10; i++{
            combo = append(combo, i)
            currentSum += i
            backtrack(combo, currentSum, i+1)
            combo = combo[:len(combo)-1]
            currentSum -= i
        }
    }
    backtrack([]int{}, 0, 1)

    return combinations
}