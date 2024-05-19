func reorganizeString(s string) string {
    counter := make(map[byte]int)
    for i := 0; i < len(s); i++{
        counter[s[i]]++
    }

    chars := []byte{}
    for k := range counter{
        chars = append(chars, k)
    }

    sort.Slice(chars, func(i,j int) bool{
        return counter[chars[i]] > counter[chars[j]]
    })

    if counter[chars[0]] > (len(s)+1)/2 {
        return ""
    }
    
    res := make([]byte, len(s))
    index := 0

    for _, c := range chars{
        for counter[c] > 0 && index < len(s){
            res[index] = c
            counter[c]--
            index+=2
        }
    }

    index = 1
    for _, c := range chars{
        for counter[c] > 0 && index < len(s){
            res[index] = c
            counter[c]--
            index+=2
        }
    }

    return string(res)
}