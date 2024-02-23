func findErrorNums(nums []int) []int {
    counter := map[int]int{}
    res := []int{} 
    counter[nums[0]] = 1
    for i := 1; i < len(nums); i++{
        counter[nums[i]]++
    }

    for i := 1; i <= len(nums); i++{
        if counter[i] == 2{
            if len(res) == 1{
                m := res[0]
                res = []int{}
                res = append(res, i)
                res = append(res, m)
                break
            }else{
                res = append(res, i)
            }
        }else if counter[i] == 0{
            res = append(res, i)
            if len(res) == 2{
                break
            }
        }
    } 
    return res

}