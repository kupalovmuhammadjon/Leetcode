func countAlternatingSubarrays(nums []int) int64 {
    var numberOfSubarrays int64
    var prev int64

    prev = 1
    numberOfSubarrays = 1
    for i := 1; i < len(nums); i++{
        if nums[i] != nums[i-1] {
            prev = prev + 1

        }else{
            prev = 1
        }
        numberOfSubarrays += prev
    }

    return numberOfSubarrays
}