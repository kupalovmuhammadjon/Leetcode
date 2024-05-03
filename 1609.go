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
func isEvenOddTree(root *TreeNode) bool {
    level := 0

    queue := Queue{}
    queue.enqueue(root)
    for queue.Length > 0{
        prev := 0
        isFirstNode := true
        currentLength := queue.Length
        for i := 0; i < currentLength; i++{
            curNode := queue.dequeue()
            if level % 2 == 0{
                if (curNode.Val % 2 == 0 || prev >= curNode.Val) && !isFirstNode{
                    return false
                }else if isFirstNode{
                    if curNode.Val % 2 == 0{
                        return false
                    }
                    isFirstNode = false
                }
                prev = curNode.Val
            }else{
                if (curNode.Val % 2 != 0 || prev <= curNode.Val) && !isFirstNode{
                    return false
                }else if isFirstNode{
                    if curNode.Val % 2 != 0{
                        return false
                    }
                    isFirstNode = false
                }
                prev = curNode.Val
            }
            if curNode.Left != nil {
                queue.enqueue(curNode.Left)
            }
            if curNode.Right != nil {
                queue.enqueue(curNode.Right)
            }
            fmt.Println(prev)
        }
        level++
    }
    return true
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