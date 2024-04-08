func reachableNodes(n int, edges [][]int, restricted []int) int {
    adj := map[int][]int{}
    rest := map[int]bool{}
    for _, v := range edges{
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    for _, v := range restricted{
        rest[v] = true
    }
    visited := map[int]bool{}
    queue := []int{0}
    numOfNodes := 0
    for len(queue) > 0{
        l := len(queue)
        for i := 0; i < l; i++{
            cur := queue[0]
            queue = queue[1:]
            if !rest[cur] && !visited[cur]{
                numOfNodes++
            }
            for _, v := range adj[cur]{
                if !rest[cur] && !visited[cur]{
                    queue = append(queue, v)
                }
            }
            visited[cur] = true
        }
    }
    return numOfNodes
}