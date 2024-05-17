/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfs func(node *TreeNode)*TreeNode
    dfs = func(node *TreeNode)*TreeNode{
        if node == nil{
            return nil
        }

        node.Left = dfs(node.Left)
        node.Right = dfs(node.Right)

        if node.Val == target && node.Left == nil && node.Right == nil{
            return nil
        }
        return node
    }

    dfs(root)
    if root.Val == target && root.Left == nil && root.Right == nil{
        return nil
    }

    return root
} 