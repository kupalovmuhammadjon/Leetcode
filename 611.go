func triangleNumber(nums []int) int {
    nums = mergeSort(nums, 0, len(nums)-1)
    numOfCombinations := 0
    for i := 0; i < len(nums); i++{
        for j := i+1; j < len(nums); j++{
            c := binarySearch(nums, j+1, nums[i]+nums[j])
            if  c != -1{
                numOfCombinations += c - j
                fmt.Println(nums[c], c - j)
            }
        }
    }

    return numOfCombinations
}

func binarySearch(nums []int, minIndex, tar int) int{
    index := -1
    left := minIndex
    right := len(nums)-1
    for left <= right{
        mid := (left + right) / 2
        if nums[mid] >= tar{
            right = mid-1
        }else {
            index = mid
            left = mid+1
        }
    }
    return index
}


func mergeSort(nums []int, s, e int) []int{
    if e - s + 1 <= 1{
        return nums[s : e+1]
    }

    m := (s + e) / 2
    left := mergeSort(nums, s, m)
    right :=mergeSort(nums, m+1, e)

    return merge(left, right)
}

func merge(left, right []int) []int{
    merged := make([]int, len(right)+len(left))

    leftIndex := 0
    rightIndex := 0
    arrayIndex := 0

    for leftIndex < len(left) && rightIndex < len(right){
        if left[leftIndex] <= right[rightIndex]{
            merged[arrayIndex] = left[leftIndex]
            leftIndex++
        }else{
            merged[arrayIndex] = right[rightIndex]
            rightIndex++
        }
        arrayIndex++
    }
    for leftIndex < len(left){
        merged[arrayIndex] = left[leftIndex]
        leftIndex++
        arrayIndex++
    }
    for rightIndex < len(right){
        merged[arrayIndex] = right[rightIndex]
        rightIndex++
        arrayIndex++
    }
    return merged
}