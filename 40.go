import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    combinations := [][]int{}

    sort.Ints(candidates) 

    var backtrack func(combo []int, currentSum int, index int)
    backtrack = func(combo []int, currentSum int, index int) {
        if currentSum == target {
            copyCombo := make([]int, len(combo))
            copy(copyCombo, combo)
            combinations = append(combinations, copyCombo)
            return
        } else if currentSum > target || index == len(candidates) {
            return
        }

        for i := index; i < len(candidates); i++ {
            if i > index && candidates[i] == candidates[i-1] {
                continue
            }

            currentSum += candidates[i]
            combo = append(combo, candidates[i])
            backtrack(combo, currentSum, i+1)
            combo = combo[:len(combo)-1]
            currentSum -= candidates[i]
        }
    }

    backtrack([]int{}, 0, 0)
    return combinations
}
