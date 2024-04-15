func networkDelayTime(times [][]int, n int, k int) int {
    graph := makeGraph(times)
    parents := make(map[int]int)
    costs := make(map[int]int)
    costs[k] = 0

    visited := make(map[int]bool)
    currentNode := findTheLowestNode(costs, visited)
    for currentNode != -1 {
        cost := costs[currentNode]

        neighbours := graph[currentNode]
        for node, time := range neighbours {
            newCost := cost + time
            if cost, ok := costs[node]; !ok || newCost < cost {
                costs[node] = newCost
                parents[node] = currentNode
            }
        }

        visited[currentNode] = true
        currentNode = findTheLowestNode(costs, visited)
    }

    // Check if the destination node is reachable from the source node
    if _, ok := costs[n]; !ok {
        return -1
    }

    // Check if there are any unreachable nodes
    for node := 1; node <= n; node++ {
        if _, ok := costs[node]; !ok {
            return -1
        }
    }

    maxTime := 0
    for _, time := range costs {
        if time > maxTime {
            maxTime = time
        }
    }
    return maxTime
}


func makeGraph(times [][]int) map[int]map[int]int {
    graph := map[int]map[int]int{}

    for _, node := range times{
        from := node[0]
        destination := node[1]
        cost := node[2]

        if _, ok := graph[from]; !ok{
            graph[from] = make(map[int]int)
        } 
        graph[from][destination] = cost
    }
    return graph
}

func findTheLowestNode(costs map[int]int, visited map[int]bool) int{
    lowestNode := -1
    lowestCostOfNode := math.MaxInt64
    for node, cost := range costs{
        if lowestCostOfNode > cost && !visited[node]{
            lowestCostOfNode = cost
            lowestNode = node
        }
    }
    return lowestNode
}