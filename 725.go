/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	arr := make([]*ListNode, k)
	runner := head
	ln := 0
	diff := 1
	q := 0
	for runner != nil {
		ln++
		runner = runner.Next
	}

	diff = ln / k
	q = ln % k

	for i := 0; i < k && head != nil; i++ {
		arr[i] = head
		size := diff
		if q > 0 {
			size++
			q--
		}
		for j := 1; j < size; j++ {
			head = head.Next
		}
		next := head.Next
		head.Next = nil
		head = next
	}

	return arr
}