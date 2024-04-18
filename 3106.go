func getSmallestString(s string, k int) string {
    newStringWithKDiff := ""
    for i := 0; i < len(s); i++{
        if s[i] == 'a'{
            newStringWithKDiff += string(s[i])
            continue
        }
        distanceToZ := 123 - int(s[i])
        distanceToA := int(s[i]) - 97
        if distanceToZ > distanceToA && k >= distanceToA{
            k -= distanceToA
            newStringWithKDiff += "a"
        }else if distanceToZ < distanceToA && k >= distanceToZ{
            k -= distanceToZ
            newStringWithKDiff += "a"
        }else if distanceToZ == distanceToA && k >= distanceToZ{
            k -= distanceToZ
            newStringWithKDiff += "a"
        }else if k > 0{
            newStringWithKDiff += string(int(s[i]) - k)
            k = 0
        }else{
            newStringWithKDiff += string(s[i])
        }
    }
    return newStringWithKDiff
}