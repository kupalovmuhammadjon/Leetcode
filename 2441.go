func findMaxK(nums []int) int {
    maxNumber := math.MinInt64
    numbers := map[int]bool{}
    for _, num := range nums{
        numbers[num] = true
        if numbers[-1*num] && maxNumber < abs(num) {
            maxNumber = abs(num)
        }
    }
    if maxNumber == math.MinInt64{
        return -1
    }
    return maxNumber
}
func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}