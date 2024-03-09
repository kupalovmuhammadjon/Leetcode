/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	for head != nil {
		if head.Val == -1234324311 {
			return true
		}
		head.Val = -1234324311
		head = head.Next
	}

	return false
}