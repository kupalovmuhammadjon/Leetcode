func separateDigits(nums []int) []int {
    res := []int{}

    for _, v := range nums{
        temp := []int{}
        for v > 0{
            temp = append(temp, v%10)
            v /= 10
        }
        for i := len(temp)-1; i >= 0; i--{
            res = append(res, temp[i])
        }
    }

    return res

}