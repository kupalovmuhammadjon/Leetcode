/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		level := []int{}
		for i := 0; i < l; i++ {
			cur := q[0]
			level = append(level, cur.Val)
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}