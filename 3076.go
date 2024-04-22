func shortestSubstrings(arr []string) []string {
    substrings := map[string]int{}
    answer := []string{}

    for _, st := range arr {
        for i := 0; i < len(st); i++ {
            for j := i + 1; j <= len(st); j++ {
                substrings[st[i:j]]++
            }
        }
    }

    for _, st := range arr {
        current := map[string]int{}
        currentAnswer := []string{}
        for i := 0; i < len(st); i++ {
            for j := i + 1; j <= len(st); j++ {
                substr := st[i:j]
                current[substr]++
                if current[substr] == substrings[substr] {
                    currentAnswer = append(currentAnswer, substr)
                }
            }
        }

        if len(currentAnswer) > 0 {
            sort.Slice(currentAnswer, func(i, j int) bool {
                if len(currentAnswer[i]) == len(currentAnswer[j]) {
                    return currentAnswer[i] < currentAnswer[j]
                }
                return len(currentAnswer[i]) < len(currentAnswer[j])
            })
            answer = append(answer, currentAnswer[0])
        } else {
            answer = append(answer, "")
        }
    }
    return answer
}
