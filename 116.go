/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
	if root == nil{
        return root
    }
	// breadth first search approach
    queue := []*Node{root}
    for len(queue) > 0{
        length := len(queue)
        leaves := []*Node{}
        for i := 0; i < length; i++{
            currentNode := queue[0]
            queue = queue[1:]
			// Populating next pointers
            if len(queue) > 0{
                currentNode.Next = queue[0]
            }else{
                currentNode.Next = nil
            }
			// adding leaves to queue
            if currentNode.Left != nil{
                leaves = append(leaves, currentNode.Left)
            }
            if currentNode.Right != nil{
                leaves = append(leaves, currentNode.Right)
            }
        }
        queue = leaves
    }
    return root
}