func makesquare(matchsticks []int) bool {
    totalLength := 0
    for _, length := range matchsticks {
        totalLength += length
    }

    if totalLength%4 != 0 {
        return false
    }

    sideLength := totalLength / 4
    sides := make([]int, 4)

    sort.Slice(matchsticks, func(i, j int) bool {
        return matchsticks[i] > matchsticks[j]
    })

    var backtrack func(ind int) bool
    backtrack = func(ind int) bool{
        if ind == len(matchsticks){
            for i := 1; i < 4; i++{
                if sides[i] != sides[i-1]{
                    return false  
                } 
            }
            return true
        }

        for i := 0; i < 4; i++{
            if sideLength >= sides[i] + matchsticks[ind]{
                sides[i] += matchsticks[ind]
                if backtrack(ind+1){
                    return true
                }
                sides[i] -= matchsticks[ind]
            }
        }
        return false 
    }

    return backtrack(0)
}