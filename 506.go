func findRelativeRanks(score []int) []string {
    answer := []string{}
    cpyScore := make([]int, len(score))
    copy(cpyScore, score)
    sort.Ints(cpyScore)
    places := map[int]int{}
    for i := 0; i < len(cpyScore); i++{
        places[cpyScore[i]] = len(cpyScore)-i
    }                  
    for _, sc := range score{
        if places[sc] > 3{
            answer = append(answer, strconv.Itoa(places[sc]))
        }else if places[sc] == 1{
            answer = append(answer, "Gold Medal")
        }else if places[sc] == 2{
            answer = append(answer, "Silver Medal")
        }else{
            answer = append(answer, "Bronze Medal")
        }
    }
    
    return answer
}