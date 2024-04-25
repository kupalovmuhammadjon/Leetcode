/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }

    answer := [][]int{}
    queue := []*TreeNode{root}
    order := true
    for len(queue) > 0{
        level := []int{}
        length := len(queue)
        leaves := []*TreeNode{}
        for i := 0; i < length; i++{
            currentNode := queue[i]
            if currentNode.Left != nil{
                leaves = append(leaves, currentNode.Left)
            }
            if currentNode.Right != nil{
                leaves = append(leaves, currentNode.Right)
            }
        }
        if order{
            for i := 0; i < length; i++{
                currentNode := queue[0]
                queue = queue[1:]
                level = append(level, currentNode.Val)
            }
            answer = append(answer, level)
            order = false
        }else{
            for i := length-1; i >= 0; i--{
                currentNode := queue[i]
                queue = queue[:i]
                level = append(level, currentNode.Val)
            }
            answer = append(answer, level)
            order = true
        }
        queue = leaves
    }
    return answer
}