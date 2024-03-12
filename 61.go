/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	runner := head
	last := dummy
	ln := 0
	for runner != nil {
		ln++
		last = runner
		runner = runner.Next
	}
	if ln < 2 || k == 0 {
		return head
	}
	if ln-(k%ln) == ln {
		return head
	}
	runner = head
	cur := 0
	save := head

	for runner != nil {
		if cur == ln-(k%ln) {
			dummy.Next = runner
			save.Next = nil
			last.Next = head
			break
		}
		save = runner
		runner = runner.Next
		cur++
	}

	return dummy.Next
}