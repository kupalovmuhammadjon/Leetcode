/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	start := &ListNode{Val: 0, Next: nil}
	end := &ListNode{Val: 0, Next: nil}
	h := start
	eh := end
	cur := 1
	for head != nil {
		if cur%2 == 1 {
			start.Next = &ListNode{Val: head.Val, Next: nil}
			start = start.Next
		} else {
			end.Next = &ListNode{Val: head.Val, Next: nil}
			end = end.Next
		}
		head = head.Next
		cur++
	}
	start.Next = eh.Next
	return h.Next
}