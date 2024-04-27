/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    depth := 1
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int{
        if root == nil {
            return math.MaxInt64
        }
        if root.Left == nil && root.Right == nil{
            return depth
        }
        return min(dfs(root.Right, depth+1), dfs(root.Left, depth+1))
    }

    return dfs(root, depth)
}