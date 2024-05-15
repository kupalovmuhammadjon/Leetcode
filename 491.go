func findSubsequences(nums []int) [][]int {
    subsequences := [][]int{}

    var backtrack func(curSub []int, ind int)
    backtrack = func(curSub []int, ind int){
        if len(curSub) >= 2{
            subsequences = append(subsequences, append([]int{}, curSub...))
        }

        used := map[int]bool{}
        for i := ind; i < len(nums); i++{
            if (len(curSub) == 0 || curSub[len(curSub)-1] <= nums[i]) && !used[nums[i]] {
                used[nums[i]] = true
                curSub = append(curSub, nums[i])
                backtrack(curSub, i+1)
                curSub = curSub[:len(curSub)-1]
            }
        }
    }
    backtrack([]int{}, 0)

    return subsequences
}