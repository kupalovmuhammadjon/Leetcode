func maxProduct(nums []int) int {
    maxPr := math.MinInt64
    for _, v := range nums{
        if maxPr < v{
            maxPr = v
        }
    }
    currentMax := 1
    currentMin := 1
    for _, n := range nums{
        if n == 0{
            currentMax, currentMin = 1, 1
            continue
        }
        currentPr := n * currentMax
        currentMax = max(currentPr, n * currentMin, n)
        currentMin = min(currentPr, n * currentMin, n)
        maxPr = max(currentMax, maxPr)
    }

    return maxPr
}