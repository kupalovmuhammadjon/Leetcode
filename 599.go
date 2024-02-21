func findRestaurant(list1 []string, list2 []string) []string {
    res := []string{}
    mn := 3459
    for _, word := range list1{
        i1 := index(list1, word)
        i2 := index(list2, word)
        if i1 != -1 && i2 != -1{
            if mn >= i1 + i2{
                mn = i1 + i2
            }
        }
    }
    for _, word := range list1{
        i1 := index(list1, word)
        i2 := index(list2, word)
        if i1 != -1 && i2 != -1{
            if mn >= i1 + i2{
                res = append(res, list1[i1])
            }
        }
    }
    return res
}
func index(ls []string, word string) int{
    for i := 0; i < len(ls); i++{
        if ls[i] == word{
            return i
        }
    }
    return -1
}