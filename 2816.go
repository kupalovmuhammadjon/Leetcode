/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func doubleIt(head *ListNode) *ListNode {
    leftOver := 0
    var double func(node *ListNode) 
    double = func(node *ListNode) {
        if node.Next == nil{
            return
        }
        node = node.Next
        double(node)
        node.Val = node.Val * 2 + leftOver
        leftOver = node.Val / 10
        node.Val = node.Val % 10
    }
    double(head)
    head.Val = head.Val * 2 + leftOver
    if head.Val > 9{
        newHead := &ListNode{head.Val/10, head}
        head.Val = head.Val % 10
        return newHead
    }
    return head    
}