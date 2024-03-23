/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {

	return countGoodNodes(root, root.Val)
}

func countGoodNodes(root *TreeNode, mx int) int {
	if root == nil {
		return 0
	}

	count := 0
	if mx <= root.Val {
		mx = root.Val
		count++
	}

	return count + countGoodNodes(root.Left, mx) + countGoodNodes(root.Right, mx)
}