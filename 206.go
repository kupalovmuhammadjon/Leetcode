/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil{
	//     return head
	// }
	// h := reverseList(head.Next)
	// head.Next.Next = head
	// head.Next = nil

	// return h

	var prev *ListNode
	cur := head

	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}
