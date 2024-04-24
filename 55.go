func canJump(nums []int) bool {
    prev := nums[0]
    for i := 1; i < len(nums); i++{
        if prev == 0{
            return false
        }
        prev = max(prev-1, nums[i])
    }
    return true
}