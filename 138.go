/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil{
        return nil
    }

    nodes := map[*Node]*Node{}
    h := head
    for h != nil{
        newNode := &Node{}
        newNode.Val = h.Val
        nodes[h] = newNode
        h = h.Next
    }
    h = head
    newList := nodes[h]
    newHead := newList
    for h != nil{
        newList.Random = nodes[h.Random]    
        h = h.Next
        newList.Next = nodes[h]
        newList = newList.Next
    }

    return newHead
}