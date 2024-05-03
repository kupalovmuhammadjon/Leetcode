/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func binaryTreePaths(root *TreeNode) []string {
    pathes := []string{}

    var dfs func (node *TreeNode, currentPath string)
    dfs = func (node *TreeNode, currentPath string){
        if node == nil{
            return
        }
        currentPath += strconv.Itoa(node.Val)
        if node.Left == nil && node.Right == nil{
            pathes = append(pathes, currentPath)
        }
        currentPath += "->"
        dfs(node.Left, currentPath)
        dfs(node.Right, currentPath)
    }
    dfs(root, "")

    return pathes
}