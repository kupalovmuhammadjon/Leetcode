type ShortestPath struct{
    Graph map[int]map[int]int
    Parents map[int]int
    Costs map[int]int
}

func costructor(edges [][]int) *ShortestPath{
    graph := map[int]map[int]int{}

    for _, edge := range edges{
        from := edge[0]
        to := edge[1]
        weight := edge[2]
        if _, ok := graph[from]; !ok{
            graph[from] = make(map[int]int)
        }
        if _, ok := graph[to]; !ok{
            graph[to] = make(map[int]int)
        }
        graph[from][to] = weight
        graph[to][from] = weight
    }
    return &ShortestPath{graph, map[int]int{}, map[int]int{}}
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    shp := costructor(edges)
    city := [2]int{-1, math.MaxInt64}
    for i := 0; i < n; i++{
        shp.Costs = map[int]int{}
        shp.Parents = map[int]int{}
        reachableCities := dijkstra(shp, i, distanceThreshold)
        if reachableCities < city[1]{
            city[1] = reachableCities
            city[0] = i
        }else if reachableCities <= city[1] && city[0] < i{
            city[1] = reachableCities
            city[0] = i
        }
    } 
    return city[0]
}

func dijkstra(shp *ShortestPath, startNode, distanceThreshold int) int{
    visited := map[int]bool{}
    shp.Costs[startNode] = 0
    currentNode := shp.FindLowestNode(visited)
    for currentNode != -1{
        currentCost := shp.Costs[currentNode]
        for node, cost := range shp.Graph[currentNode]{
            newCost := currentCost + cost
            if _, ok := shp.Costs[node]; !ok || shp.Costs[node] > newCost{
                shp.Costs[node] = newCost
                shp.Parents[node] = currentNode
            }
        }
        visited[currentNode] = true
        currentNode = shp.FindLowestNode(visited)
    }
    
    numberOfReachableCities := 0
    for _, cost := range shp.Costs{
        if cost <= distanceThreshold{
            numberOfReachableCities++
        }
    }
    return numberOfReachableCities
}

func (shp *ShortestPath) FindLowestNode(visited map[int]bool) int{
    lowestNode := -1
    lowestCost := math.MaxInt64

    for node, cost := range shp.Costs{
        if lowestCost > cost && !visited[node]{
            lowestCost = cost 
            lowestNode = node
        }
    }

    return lowestNode
}