import (
	"math"
)
func findPairs(nums []int, k int) int {
	count := 0
	pairs := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i] - nums[j]))) == k {
				pair := []int{nums[i], nums[j]}
				if notInit(pairs, pair) {
					count++
					pairs = append(pairs, pair)
					fmt.Println(nums[i], nums[j])
				}
			}
		}
	}
	return count
}

func notInit(ls [][]int, pair []int) bool{
    for _, p := range ls{
        if p[0] == pair[0] && p[1] == pair[1] || p[0] == pair[1] && p[1] == pair[0]{
            return false
        }
    }
    return true
}