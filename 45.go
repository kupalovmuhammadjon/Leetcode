func jump(nums []int) int {
    numberOfJumps := 0
    ln := len(nums) - 1
    for ln != 0{
        for i := 0; i < ln; i++{
            if i + nums[i] >= ln{
                ln = i
                numberOfJumps++
                break
            }
        } 
    }
    return numberOfJumps
}