/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	maxOfEachLevel := []int{}
	if root == nil {
		return maxOfEachLevel
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		mx := math.MinInt
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if mx < cur.Val {
				mx = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		maxOfEachLevel = append(maxOfEachLevel, mx)
	}
	return maxOfEachLevel
}