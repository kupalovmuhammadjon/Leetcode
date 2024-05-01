func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++{
        j := binarySearch(numbers, target - numbers[i], i)
        if j != -1{
            return []int{i+1, j+1}
        }
    }
    return []int{}
}
func binarySearch(nums []int, tar, restr int) int{
    l := 0
    r := len(nums)-1

    for l <= r{
        m := (l+r) / 2
        if nums[m] == tar && restr != m{
            return m
        }else if nums[m] > tar{
            r = m - 1
        }else{
            l = m + 1
        }
    }
    return -1
}