/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	count := 0
	mp := map[int]int{}
	mp[0] = 1

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}

		curSum += node.Val
		count += mp[curSum-targetSum]
		mp[curSum]++

		dfs(node.Left, curSum)
		dfs(node.Right, curSum)

		mp[curSum]--
	}

	dfs(root, 0)

	return count
}
