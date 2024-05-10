/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    preoder := []int{}
    if root == nil{
        return preoder
    }

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }

        preoder = append(preoder, node.Val)
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)

    return preoder
}