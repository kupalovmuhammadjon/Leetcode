func findPeaks(mountain []int) []int {
    peaks := []int{}

    for i := 1; i < len(mountain)-1;i++{
        if mountain[i-1] < mountain[i] && mountain[i+1] < mountain[i]{
            peaks = append(peaks, i)
            
        }
    }
    return peaks
}