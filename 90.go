func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)

    subsets := [][]int{}
    var backtrack func(currentSet []int, index int)
    backtrack = func(currentSet []int, index int){

        copySet := make([]int, len(currentSet))
        copy(copySet, currentSet)
        subsets = append(subsets, copySet)
        for i := index; i < len(nums); i++{
            if i > index && nums[i] == nums[i-1]{
                continue
            }
            currentSet = append(currentSet, nums[i])
            backtrack(currentSet, i+1)
            currentSet = currentSet[:len(currentSet)-1]

        }
    }
    backtrack([]int{}, 0)

    return subsets
}