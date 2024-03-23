/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	r1 := searchBST(root.Left, val)
	r2 := searchBST(root.Right, val)
	if r1 != nil {
		return r1
	}
	if r2 != nil {
		return r2
	}
	return nil
}