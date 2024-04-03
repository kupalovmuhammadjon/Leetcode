/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		temp := []int{}
		for i := 0; i < l; i++ {
			cur := q[0]
			temp = append(temp, cur.Val)
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append([][]int{temp}, res...)
	}
	return res
}