func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0{
        return true
    }
    adj := map[int][]int{}
    for _, v := range prerequisites{
        adj[v[0]] = append(adj[v[0]], v[1])
    }

    visited := map[int]bool{}
    for k := range adj{
        if !dfs(adj, visited, k) {
            return false
        }
    }
    return true
}

func dfs(adj map[int][]int, visited map[int]bool, node int) bool{
    if visited[node]{
        return false
    }
    if len(adj[node]) == 0{
        return true
    }

    visited[node] = true
    for i := 0; i < len(adj[node]); i++{
        if !dfs(adj, visited, adj[node][i]){
            return false
        }
    }
    visited[node] = false
    adj[node] = []int{}
    
    return true 
}