func largestNumber(nums []int) string {
    var number strings.Builder
    strNums := make([]string, len(nums))
    for i, num := range nums{
        strNums[i] = strconv.Itoa(num)
    }
    
    sort.Slice(strNums, func(i, j int) bool{
        return strNums[i] + strNums[j] > strNums[j] + strNums[i]
    })

    for _, num := range strNums{
        number.WriteString(num)
    }

    if number.String()[0] == '0'{
        return "0"
    }
    
    return number.String()
}