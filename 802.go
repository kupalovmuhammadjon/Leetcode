func eventualSafeNodes(graph [][]int) []int {
    safeNodes := []int{}

    restricted := map[int]bool{}
    safe := map[int]bool{}
    var dfs func(curNode int, visited map[int]bool) bool
    dfs = func(curNode int, visited map[int]bool) bool{
        if visited[curNode] || restricted[curNode]{
            return false
        }else if safe[curNode]{
            return true
        }
        visited[curNode] = true
        for i := 0; i < len(graph[curNode]); i++{
            isSafe := dfs(graph[curNode][i], visited)
            if !isSafe{
                restricted[graph[curNode][i]] = true
                return false
            }else{
                safe[graph[curNode][i]] = true
            }
        }
        visited[curNode] = false
        safe[curNode] = true

        return true
    }
    // dfs(1, map[int]bool{})
    for i := 0; i < len(graph); i++{
        isSafe := dfs(i, map[int]bool{})
        if isSafe{
            safeNodes = append(safeNodes, i)
        }else{
            restricted[i] = true
        }
    }
    return safeNodes
}