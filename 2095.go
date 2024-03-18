/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	fast := head
	slow := head
	back := head
	ln := 0
	for fast.Next != nil && fast.Next.Next != nil {
		back = slow
		slow = slow.Next
		fast = fast.Next.Next
		ln += 2
	}
	for fast != nil {
		fast = fast.Next
		ln += 1
	}
	if ln == 1 {
		return nil
	} else if ln%2 == 0 {
		slow.Next = slow.Next.Next
	} else {
		back.Next = back.Next.Next
	}

	return head
}