func isPossibleDivide(nums []int, k int) bool {
    sort.Ints(nums)
    count := map[int]int{}
    for _, num := range nums{
        count[num]++
    }
    for i := 0; i < len(nums); i++{
        if count[nums[i]] > 0{
            count[nums[i]]--
            prev := nums[i]
            ln := 1
            for j := i; j < len(nums); j++{
                if ln == k{
                    break
                }
                if prev+1 == nums[j] && count[nums[j]] > 0{
                    count[nums[j]]--
                    prev = nums[j]
                    ln++
                }
            }
            if ln != k{
                return false
            }

        }
    }
    return true
}