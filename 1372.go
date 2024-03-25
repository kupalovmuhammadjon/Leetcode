/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	mx := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left := maxZigZag(cur, true)
		right := maxZigZag(cur, false)

		mx = max(mx, max(left, right))

		if cur.Right != nil && !(cur.Right.Right == nil && cur.Right.Left == nil) {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil && !(cur.Left.Right == nil && cur.Left.Left == nil) {
			stack = append(stack, cur.Left)
		}

	}
	return mx
}

func maxZigZag(node *TreeNode, direction bool) int {
	ln := 0
	for node != nil {
		if direction {
			direction = false
			node = node.Left
		} else {
			direction = true
			node = node.Right
		}
		if node != nil {
			ln++
		}
	}
	return ln
}