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