/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type LinkedList struct{
    Val *TreeNode
    Next *LinkedList
}
type Queue struct{
    Head *LinkedList
    Tail *LinkedList
    Length int
}

func (q *Queue) enqueue(val *TreeNode){
    q.Length++
    newNode := &LinkedList{val, nil}
    if q.Head == nil{
        q.Head = newNode
        q.Tail = newNode
        return 
    } 
    q.Tail.Next = newNode
    q.Tail = q.Tail.Next
}

func (q *Queue) dequeue() *TreeNode{
    q.Length--
    currentNode := q.Head
    q.Head = q.Head.Next
    if q.Head == nil{
        q.Tail = nil
        return currentNode.Val
    }
    
    return currentNode.Val
}
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    levelSums := []int{}

    queue := Queue{}
    queue.enqueue(root)

    for queue.Length > 0{
        levelSum := 0
        currentLength := queue.Length
        for i := 0; i < currentLength; i++{
            currentNode := queue.dequeue()
            levelSum += currentNode.Val
            if currentNode.Left != nil{
                queue.enqueue(currentNode.Left)
            }
            if currentNode.Right != nil{
                queue.enqueue(currentNode.Right)
            }
        }
        levelSums = append(levelSums, levelSum)
    }
    if len(levelSums) < k{
        return -1
    }
    sort.Ints(levelSums)
    fmt.Println(levelSums)
    return int64(levelSums[len(levelSums)-k])
}