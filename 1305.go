/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    slc1 := []int{}
    slc2 := []int{}
    
    dfs(root1, &slc1)
    dfs(root2, &slc2)
    
    return merge(slc1, slc2)
}

func dfs(node *TreeNode, slc *[]int) {
    if node == nil {
        return 
    }

    dfs(node.Left, slc)
    *slc = append(*slc, node.Val)
    dfs(node.Right, slc)
}

func merge(left, right []int) []int{
    result := make([]int, len(left) + len(right))

    leftIndex := 0
    rightIndex := 0
    arrayIndex := 0

    for leftIndex < len(left) && rightIndex < len(right){
        if left[leftIndex] <= right[rightIndex]{
            result[arrayIndex] = left[leftIndex]
            leftIndex++
        }else {
            result[rrayIndex] = right[rightIndex]
            rightIndex++
        }
        arrayIndex++
    }
    for leftIndex < len(left){
        result[arrayIndex] = left[leftIndex]
        leftIndex++
        arrayIndex++
    }
    for rightIndex < len(right){
        result[arrayIndex] = right[rightIndex]
        rightIndex++
        arrayIndex++
    }
    return result
}