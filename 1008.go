/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstFromPreorder(preorder []int) *TreeNode {
    root := &TreeNode{Val: preorder[0]}
    var add func(root *TreeNode, node int) *TreeNode
    add = func(root *TreeNode, node int) *TreeNode{
        if root == nil{
            return &TreeNode{Val: node}
        }
        if root.Val > node{
            root.Left = add(root.Left, node)
        }
        if root.Val < node{
            root.Right = add(root.Right, node)
        }
        return root
    }
    for i := 1; i < len(preorder); i++{
        add(root, preorder[i])
    }   
    return root
}