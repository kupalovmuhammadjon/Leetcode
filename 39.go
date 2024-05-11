func combinationSum(candidates []int, target int) [][]int {
    combinations := [][]int{}

    var backtrack func(currentCombo []int, currentSum int, indx int)
    backtrack = func(currentCombo []int, currentSum int, indx int){
        if currentSum == target{
            copyCombo := make([]int, len(currentCombo))
            copy(copyCombo, currentCombo)
            combinations = append(combinations, copyCombo)
            return
        }else if currentSum > target {
            return
        }
        for i := indx; i < len(candidates); i++{
            currentCombo = append(currentCombo, candidates[i])
            backtrack(currentCombo, currentSum + candidates[i], i)
            currentCombo = currentCombo[:len(currentCombo)-1]

        }
    }
    backtrack([]int{}, 0, 0)
    return combinations
}