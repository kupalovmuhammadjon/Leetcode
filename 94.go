/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    inorder := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }
        dfs(node.Left)
        inorder = append(inorder, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return inorder
}