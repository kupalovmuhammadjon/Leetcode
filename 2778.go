func sumOfSquares(nums []int) int {
    n := len(nums)
    sm := 0
    for i := 0; i < n; i++{
        if n % (i+1) == 0{
            sm += nums[i] * nums[i] 
        }
    }
    return sm
}