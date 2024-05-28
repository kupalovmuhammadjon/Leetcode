func searchMatrix(matrix [][]int, target int) bool {
    binarySearch := func(arr []int) bool{
        l := 0
        r := len(arr)-1

        for l <= r{
            mid := (l + r) / 2

            if arr[mid] == target{
                return true
            }
            if arr[mid] > target{
                r = mid - 1
            }else{
                l = mid + 1
            }
        }
        return false
    }

    for _, arr := range matrix{
        fmt.Println(arr)
        if binarySearch(arr){
            return true
        }
    }
    return false
}
