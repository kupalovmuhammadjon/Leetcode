/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfs func(root *TreeNode, currentSum int) bool
    dfs = func(root *TreeNode, currentSum int) bool{
        if root == nil{
            return false
        }
        currentSum += root.Val
        if currentSum == targetSum && root.Left == nil && root.Right == nil {
            return true
        }

        return dfs(root.Left, currentSum) || dfs(root.Right, currentSum)
    }
    return dfs(root, 0)
}