func maxFrequencyElements(nums []int) int {
    mxcount := 0

    for i := 0; i < len(nums); i++ {
        count := 1
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                count++
            }
        }

        if count > mxcount {
            mxcount = count
        }
    }
    res := 0
    for i := 0; i < len(nums); i++ {
        count := 1
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                count++
            }
        }

        if count == mxcount {
            res += count
        }
    }

    return res
}