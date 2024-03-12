/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	runner := head
	vals := []int{}
	for runner != nil {
		vals = append(vals, runner.Val)

		runner = runner.Next
	}
	sort.Ints(vals)
	runner = head
	i := 0
	for runner != nil {
		runner.Val = vals[i]
		i++
		runner = runner.Next
	}

	return head
}
