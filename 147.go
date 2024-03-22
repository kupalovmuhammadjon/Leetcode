/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := head

	for curr != nil {
		next := curr.Next
		prev := dummy
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}
		curr.Next = prev.Next
		prev.Next = curr
		curr = next
	}

	return dummy.Next
}