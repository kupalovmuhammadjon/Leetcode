/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapNodes(head *ListNode, k int) *ListNode {
    nodes :=  map[int]*ListNode{}
    runner := head
    ln := 0
    for runner != nil{
        ln++
        nodes[ln] = runner 
        runner = runner.Next
    }
    
    nodes[k].Val, nodes[len(nodes)-k+1].Val = nodes[len(nodes)-k+1].Val, nodes[k].Val
    return head
}