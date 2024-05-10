func kthSmallestPrimeFraction(arr []int, k int) []int {
    fractions := [][]float64{}
    for i := 0; i < len(arr); i++{
        for j := i+1; j < len(arr); j++{
            fractions = append(fractions, []float64{float64(arr[i]), float64(arr[j])})
        }
    }
    sort.Slice(fractions, func(i, j int) bool {
        return fractions[i][0]/fractions[i][1] < fractions[j][0]/fractions[j][1]
    })

    return []int{int(fractions[k-1][0]), int(fractions[k-1][1])}
}