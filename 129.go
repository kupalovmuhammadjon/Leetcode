/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumNumbers(root *TreeNode) int {
    // Depth first search approach using recursion
    return dfs(root, 0) 
}
func dfs(node *TreeNode, number int) int{
    if node == nil{
        return 0
    }
	// building up number until it hit to base case
    number = number * 10 + node.Val

    if node.Left == nil && node.Right == nil{
        return number
    }

    return dfs(node.Left, number) + dfs(node.Right, number)
}