/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxAncestorDiff(root *TreeNode) int {
    if root == nil {
        return 0
    }
    difference := 0

    var dfs func(node *TreeNode, maxSoFar, minSoFar int)
    dfs = func(node *TreeNode, maxSoFar, minSoFar int){
        if node == nil{
            return
        }
        difference = max(difference, abs(maxSoFar - node.Val), abs(minSoFar - node.Val))

        maxSoFar = max(maxSoFar, node.Val)
        minSoFar = min(minSoFar, node.Val)

        dfs(node.Left, maxSoFar, minSoFar)
        dfs(node.Right, maxSoFar, minSoFar)
    }
    dfs(root, root.Val, root.Val)

    return difference
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}