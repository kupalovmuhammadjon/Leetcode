func differenceOfSum(nums []int) int {
    sum := 0
    digits_sum := 0
    for i := 0; i < len(nums); i++{
        sum += nums[i]
        n := nums[i]
        for n > 0{
            digits_sum += n % 10
            n /= 10
        }
    }

    return sum - digits_sum
}