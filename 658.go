func findClosestElements(arr []int, k int, x int) []int {
    ans := []int{}
    closest := 0
    for i := 0; i < len(arr); i++{
        if abs(arr[i] - x) < abs(arr[closest] - x) || abs(arr[i] - x) == abs(arr[closest] - x) && arr[i] < arr[closest]{
            closest = i
            if arr[i] - x == 0{
                break
            }
        }
    }
    ans = append(ans, arr[closest])
    l := closest - 1
    r := closest + 1
    for l >= 0 && r < len(arr) && len(ans) < k{
        if abs(arr[l] - x) < abs(arr[r] - x) || abs(arr[l] - x) == abs(arr[r] - x) && arr[l] < arr[r]{
            ans = append(ans, arr[l])
            l--
        }else{
            ans = append(ans, arr[r])
            r++
        }
    }
    for len(ans) < k && l >= 0{
        ans = append(ans, arr[l])
        l--
    }
    for len(ans) < k && r < len(arr){
        ans = append(ans, arr[r])
        r++
    }
    sort.Ints(ans)
    return ans
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}