/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1{
        newRoot := &TreeNode{val, nil, nil}
        newRoot.Left = root
        return newRoot
    }
    levelsOfTree := []*TreeNode{root}
    currentDepth := 1
    rootHolder := root
    for len(levelsOfTree) > 0 {
        if currentDepth == depth-1{
            break
        }
        lengthOfLevel := len(levelsOfTree)
        for i := 0; i < lengthOfLevel; i++{
            currentNode := levelsOfTree[0]
            levelsOfTree = levelsOfTree[1:]
            if currentNode.Right != nil {
                levelsOfTree = append(levelsOfTree, currentNode.Right)
            }
            if currentNode.Left != nil {
                levelsOfTree = append(levelsOfTree, currentNode.Left)
            }
        }
        currentDepth++
    }
    for _, node := range levelsOfTree{
        node.Right = &TreeNode{val, nil, node.Right}
        node.Left = &TreeNode{val, node.Left, nil}
    }
    return rootHolder
}