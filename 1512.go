func numIdenticalPairs(nums []int) int {
    goodPairs := 0
    counts := map[int][]int{}

    for i := 0; i < len(nums); i++{
        counts[nums[i]] = append(counts[nums[i]], i)
    }

    for i, num := range nums{
        for _, ind := range counts[num]{
            if ind > i{
                goodPairs++
            }
        }
    }

    return goodPairs
}