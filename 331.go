func isValidSerialization(preorder string) bool {
    if preorder[0] == '#'{
        return false
    }
    nodes := strings.Split(preorder, ",")
    place := 1

    for i := 0; i < len(nodes); i++{
        if place <= 0{
            return false
        } 
        if nodes[i] == "#"{
            place--
        }else{
            place++
        }
    }

    return place == 0
}