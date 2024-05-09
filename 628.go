func maximumProduct(nums []int) int {
    nums = mergeSort(nums, 0, len(nums)-1)
    fmt.Println(nums)
    maxPr := max(nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3],
    nums[len(nums)-1] * nums[0] * nums[1])
    return maxPr
}

func mergeSort(nums []int, s, e int) []int {
    if e - s + 1 <= 1{
        return nums[s:e+1]
    }

    m := (s + e) / 2
    left := mergeSort(nums, s, m)
    right := mergeSort(nums, m+1, e)

    return merge(left, right)
}

func merge(left, right []int) []int {
    merged := make([]int, len(left)+len(right))

    leftIndex := 0
    rightIndex := 0
    arrayIndex := 0

    for leftIndex < len(left) && rightIndex < len(right) {
        if left[leftIndex] <= right[rightIndex] {
            merged[arrayIndex] = left[leftIndex]
            leftIndex++
        } else {
            merged[arrayIndex] = right[rightIndex]
            rightIndex++
        }
        arrayIndex++
    }

    for leftIndex < len(left) {
        merged[arrayIndex] = left[leftIndex]
        leftIndex++
        arrayIndex++
    }

    for rightIndex < len(right) {
        merged[arrayIndex] = right[rightIndex]
        rightIndex++
        arrayIndex++
    }

    return merged
}
