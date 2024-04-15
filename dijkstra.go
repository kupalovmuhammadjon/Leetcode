package main

import (
	"fmt"
	"math"
)

func main() {

	wgrph := [][]int{{1, 3, 2}, {1, 1, 3}, {1, 2, 4}, {2, 1, 4}, {3, 3, 6}, {2, 4, 5}, {4, 2, 5}, {5, 1, 6},
		{6, 4, 7}, {5, 2, 9}, {9, 1, 8}, {7, 3, 8}}

	mgrph := map[int]map[int]int{}
	for _, node := range wgrph {
		from := node[0]
		cost := node[1]
		to := node[2]
		// Initialize the map if it doesn't exist
        if _, ok := mgrph[from]; !ok {
            mgrph[from] = make(map[int]int)
        }
        if _, ok := mgrph[to]; !ok {
            mgrph[to] = make(map[int]int)
        }
		mgrph[from][to] = cost
		mgrph[to][from] = cost
	}
	fmt.Println(mgrph)

}
func shortestPath(graph map[int]map[int]int) int {
	visited := map[int]bool{}
	stack := []int{1}
	parents := map[int][]int{}
	parents[1] = append(parents[1], []int{0})

	for len(stack) > 0{
		l := len(stack)
		cur := stack[0]
		stack = stack[:len(stack)-1]
		for i := 0; i < l; i++{ 
			for k := range graph[cur]{
				
			}
			cost := graph[cur][0]
			stack = append(stack, graph[cur][1])
		}
	}

}
func shortestNode(parents map[int][]int, visited map[int]bool) int{
	var minParent int
	minValue := math.MaxInt64
	for k := range parents{
		if !visited[k]{
			if minValue > parents[k][0]{
				minValue = parents[k][0]
				minParent = k
			}
		}
	}
	return minParent
}