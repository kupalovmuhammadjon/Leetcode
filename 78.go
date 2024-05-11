func subsets(nums []int) [][]int {
    subsets := [][]int{}

    var backtrack func(currentSubset []int, index int)
    backtrack = func(currentSubset []int, index int){
        if len(currentSubset) == len(nums){
            return
        }
        copySubset := make([]int, len(currentSubset))
        copy(copySubset, currentSubset)
        fmt.Println(currentSubset)
        subsets = append(subsets, copySubset)
        fmt.Println(index)
        for i := index; i < len(nums); i++{
            currentSubset = append(currentSubset, nums[i])
            backtrack(currentSubset, i+1)
            currentSubset = currentSubset[:len(currentSubset)-1]
        }
    } 
    backtrack([]int{}, 0)
    subsets = append(subsets, nums)
    return subsets
}