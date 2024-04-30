package main

import "fmt"

// type LinkedList struct {
// 	val  string
// 	next *LinkedList
// }

// type Queue struct {
// 	head      *LinkedList
// 	connector *LinkedList
// 	tail      *LinkedList
// 	len       int
// }

func main() {
	graph := map[string][]string{}
	graph["s"] = []string{"1", "2"}
	graph["1"] = []string{"3", "f"}
	graph["2"] = []string{"3", "4"}
	graph["4"] = []string{"f"}        

	visited := map[string]bool{}

	queue := []string{"s"}
	distance := 0
	for len(queue) > 0{
		l := len(queue)
		distance++   
		fmt.Println(queue)                   
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur != "f"{
				visited[cur] = true
			}
			if cur == "f" {
				fmt.Println("Topildi ", distance)
			}
			for _, v := range graph[cur] {
				if !visited[v]{ 
					queue = append(queue, v)
				}
			}
		}
	}

}

// queue := Queue{head: &LinkedList{"anuj", nil}, tail: &LinkedList{"anuj", nil}, connector: &LinkedList{"anuj", nil}, len: 1}
// 	l := queue.len
// 	for i := 0; i < l; i++ {
// 		cur := queue.head
// 		if !visited[cur.val] {
// 			if cur.val == "kosheng" {
// 				fmt.Println("Topildi toxta")
// 				break
// 			}
// 			visited[cur.val] = true
// 		}
// 		for _, v := range graph[cur.val] {
// 			if !visited[v] {
// 				newTail := &LinkedList{v, nil}
// 				queue.tail.next = newTail
// 				queue.tail = queue.tail.next
// 			}
// 		}
// 	}
