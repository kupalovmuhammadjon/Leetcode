/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func nextLargerNodes(head *ListNode) []int {
    res := []int{}
    for head != nil{
        runner := head.Next
        blen := len(res)
        for runner != nil{
            if runner.Val > head.Val{
                res = append(res, runner.Val)
                break
            }
            runner = runner.Next
        }
        if len(res) == blen{
            res = append(res, 0)
        }
        head = head.Next
    }
    return res
}