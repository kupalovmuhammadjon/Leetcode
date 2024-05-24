func allPathsSourceTarget(graph [][]int) [][]int {
    paths := [][]int{}
    target := len(graph) - 1  
    var dfs func(currentPath []int, curNode int)
    dfs = func(currentPath []int, curNode int){
        if len(currentPath) > 0 && currentPath[len(currentPath)-1] == target{
            paths = append(paths, append([]int{}, currentPath...))
            return
        }

        for i := 0; i < len(graph[curNode]); i++{
            currentPath = append(currentPath, graph[curNode][i])
            dfs(currentPath, graph[curNode][i])
            currentPath = currentPath[:len(currentPath)-1]
        }
    }
    dfs([]int{0}, 0)

    return paths
}