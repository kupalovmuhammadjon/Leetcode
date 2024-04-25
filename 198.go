func rob(nums []int) int {
    fmt.Pr
    firstMove := 0
    secondMove := 0
    for i := 0; i < len(nums); i++{
        t := max(nums[i] + firstMove, secondMove)
        firstMove = secondMove
        secondMove = t
    }
    return secondMove
}