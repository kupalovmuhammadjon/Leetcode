/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	var prev *ListNode
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	first := head
	second := prev
	mx := 0
	for first != nil && second != nil {
		if mx < first.Val+second.Val {
			mx = first.Val + second.Val
		}
		first = first.Next
		second = second.Next
	}
	return mx
}