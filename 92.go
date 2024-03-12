/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	group := &ListNode{Val: 0, Next: nil}
	var end *ListNode
	var prv *ListNode
	cur := left
	lamp := true
	runner := dummy

	for i := 0; i < left-1; i++ {
		runner = runner.Next
	}

	prv = runner
	end = runner.Next

	for cur <= right && runner.Next != nil {
		if cur <= right {
			tp := &ListNode{Val: runner.Next.Val, Next: group.Next}
			group.Next = tp
			if lamp {
				end = group.Next
				lamp = false
			}
		} else if cur > right {
			break
		}
		runner = runner.Next
		cur++
	}
	prv.Next = group.Next
	end.Next = runner.Next

	return dummy.Next
}
