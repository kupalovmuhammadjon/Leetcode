func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))

    var visit func(keys []int)
    visit = func(keys []int){
        for _, key := range keys{
            if visited[key] {
                continue
            }
            visited[key] = true
            visit(rooms[key])
        }
    }
    visit([]int{0})

    for _, v := range visited{
        if !v{
            return false
        }
    }
    return true
}
