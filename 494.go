func findTargetSumWays(nums []int, target int) int {
    numberOfWays := 0

    var backtrack func(currentSum int, i int)
    backtrack = func(currentSum int, i int){
        if i == len(nums) && currentSum == target{
            numberOfWays++
            return
        }else if i == len(nums){
            return
        }

        backtrack(currentSum + (-1*nums[i]), i+1)
        backtrack(currentSum + nums[i], i+1)
    }
    backtrack(0, 0)

    return numberOfWays
}