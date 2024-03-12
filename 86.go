/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	start := &ListNode{Val: 0, Next: nil}
	end := &ListNode{Val: 0, Next: nil}
	h := start
	eh := end

	for head != nil {
		if head.Val < x {
			start.Next = &ListNode{Val: head.Val, Next: nil}
			start = start.Next
		} else {
			end.Next = &ListNode{Val: head.Val, Next: nil}
			end = end.Next
		}
		head = head.Next
	}
	start.Next = eh.Next
	return h.Next
}