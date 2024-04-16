package main

import (
	"fmt"
	"math"
)

func main() {

	wgrph := [][]int{{1, 3, 2}, {1, 1, 3}, {1, 2, 4}, {2, 1, 4}, {3, 3, 6}, {2, 4, 5}, {4, 2, 5}, {5, 1, 6},
		{6, 4, 7}, {5, 2, 9}, {9, 1, 8}, {7, 3, 8}}

	graph := createDirectionOfRoutes(wgrph)

	costs := map[int]int{}
	costs[1] = 0

	parents := map[int]int{}
	visited := map[int]bool{}

	node := findLowestCostNode(costs, visited)
	for node != -1 {
		cost := costs[node]
		neighbours := graph[node]
		for k := range neighbours {
			newCost := cost + neighbours[k]
			if _, ok := costs[k]; !ok || costs[k] > newCost {
				costs[k] = newCost
				parents[k] = node
			}
		}
		visited[node] = true
		node = findLowestCostNode(costs, visited)
	}
	fmt.Println(costs[7])
	fmt.Println(parents)
	path := []int{}
	city := 7
	for city != 1{
		if city == 1{
			path = append(path, city)
			break
		}
		path = append(path, city)
		city = parents[city]
	}
	fmt.Println(path)
}
func createDirectionOfRoutes(mgraph [][]int) map[int]map[int]int {
	graph := map[int]map[int]int{}

	for _, node := range mgraph {
		from := node[0]
		cost := node[1]
		destination := node[2]

		if _, ok := graph[from]; !ok {
			graph[from] = make(map[int]int)
		}
		graph[from][destination] = cost
	}
	return graph
}

func findLowestCostNode(costs map[int]int, visited map[int]bool) int {
	lowestCost := math.MaxInt64
	lowestCostNode := -1
	for k := range costs {
		cost := costs[k]
		if lowestCost > cost && !visited[k] {
			lowestCost = cost
			lowestCostNode = k
		}
	}

	return lowestCostNode
}
