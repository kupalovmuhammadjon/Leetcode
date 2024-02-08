func maximizeSum(nums []int, k int) int {
    sum := 0
    mx := 0
    for _, v := range nums{
        
        if mx < v{
            mx = v
        }
    }
    for i := 0; i < k; i++{
        sum += mx
        mx++
    }
    return sum
}