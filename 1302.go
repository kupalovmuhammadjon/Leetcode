/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deepestLeavesSum(root *TreeNode) int {
    if root == nil{
        return 0
    }

    sumOfLeaves := 0
    queue := []*TreeNode{root}
    for len(queue) > 0{
        length := len(queue)
        sumOfLeaves = 0
        for i := 0; i < length; i++{
            currentNode := queue[0]
            sumOfLeaves += currentNode.Val
            queue = queue[1:]
            if currentNode.Left != nil{
                queue = append(queue, currentNode.Left)
            }
            if currentNode.Right != nil{
                queue = append(queue, currentNode.Right)
            }
        }
    }

    return sumOfLeaves
}