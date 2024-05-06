/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNodes(head *ListNode) *ListNode {
    maxValueFromBack := math.MinInt64
    var last *ListNode
    var reverseIteration func(currentNode *ListNode) 
    reverseIteration = func(currentNode *ListNode) {
        if currentNode.Next == nil{
            return
        }
        currentNode = currentNode.Next
        reverseIteration(currentNode)
        if currentNode.Val >= maxValueFromBack{
            currentNode.Next = last
            last = currentNode
            maxValueFromBack = currentNode.Val
        }
    }
    reverseIteration(head)
    if head.Val >= maxValueFromBack{
        head.Next = last
        return head
    }
    return last
}

