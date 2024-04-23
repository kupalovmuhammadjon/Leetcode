/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) [][]int {
    answer := [][]int{}
    var dfs func(*TreeNode, []int, int)
	// depth first search appproach
    dfs = func(node *TreeNode, path []int, currentSum int){
        if node == nil {
            return 
        }
        currentSum += node.Val
        path = append(path, node.Val)
        if node.Left == nil && node.Right == nil && currentSum == targetSum {
            newPath := make([]int, len(path))
            copy(newPath, path)
            answer = append(answer, newPath)
        }

        dfs(node.Left, path, currentSum)
        dfs(node.Right, path, currentSum)
    }
    dfs(root, []int{}, 0)

    return answer
}