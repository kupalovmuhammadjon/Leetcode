/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the list
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the list
	prev.Next = nil
	var secondHead *ListNode
	cur := slow
	for cur != nil {
		nxt := cur.Next
		cur.Next = secondHead
		secondHead = cur
		cur = nxt
	}

	// Merge the two halves
	cur = head
	last := cur
	for secondHead != nil {
		if cur == nil {
			last.Next = secondHead
			last = last.Next
			last.Next = nil
			break
		}
		nxt := cur.Next
		cur.Next = secondHead
		secondHead = secondHead.Next
		cur.Next.Next = nxt
		last = cur.Next
		cur = nxt
	}

}
