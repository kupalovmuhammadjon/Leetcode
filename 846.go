func isNStraightHand(hand []int, groupSize int) bool {
    sort.Ints(hand)
    count := map[int]int{}
    for _, num := range hand{
        count[num]++
    }
    for i := 0; i < len(hand); i++{
        if count[hand[i]] > 0{
            count[hand[i]]--
            prev := hand[i]
            ln := 1
            for j := i; j < len(hand); j++{
                if ln == groupSize{
                    break
                }
                if prev+1 == hand[j] && count[hand[j]] > 0{
                    count[hand[j]]--
                    prev = hand[j]
                    ln++
                }
            }
            if ln != groupSize{
                return false
            }

        }
    }
    return true
}