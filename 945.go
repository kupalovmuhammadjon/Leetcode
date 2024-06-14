func minIncrementForUnique(nums []int) int {
    sort.Ints(nums)
    numberOfIncrements := 0
    exist := map[int]bool{}
    exist[nums[0]] = true
    for i := 1; i < len(nums); i++{
        if exist[nums[i]]{
            diff := nums[i-1]-nums[i]+1
            nums[i] += diff
            numberOfIncrements += diff
        }
        exist[nums[i]] = true
    }
    return numberOfIncrements
}