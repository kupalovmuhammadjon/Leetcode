/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthSmallest(root *TreeNode, k int) int {
    ans := -1   
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil{
            return 
        }
        dfs(node.Left)
        if k > 0{
            k--
            if k == 0{
                ans = node.Val
            }
        }
        dfs(node.Right)
    }
    dfs(root)

    return ans
}

