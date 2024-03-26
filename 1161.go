/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	maxSum := math.MinInt64
	level := 1

	q := []*TreeNode{root}
	treeLevel := 1
	for len(q) > 0 {
		sm := 0
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			sm += cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		if maxSum < sm {
			maxSum = sm
			level = treeLevel
		}
		treeLevel++
	}
	return level
}