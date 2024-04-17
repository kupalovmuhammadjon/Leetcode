func minOperationsToMakeMedianK(nums []int, k int) int64 {
    sort.Ints(nums)
    middle := len(nums) / 2
    
    var numberOfOpertations int64
    if nums[middle] > k{
        for i := middle; i >= 0 && nums[i] > k; i--{
            numberOfOpertations += int64(nums[i] - k)
        }
    }else{
        for i := middle; i < len(nums) && nums[i] < k; i++{
            numberOfOpertations += int64(k - nums[i])
        }
    }
    return numberOfOpertations
}