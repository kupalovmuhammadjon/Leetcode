/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
    inorderTraversal []int
}


func Constructor(root *TreeNode) BSTIterator {
    inorderTraversal := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }

        dfs(node.Left)
        inorderTraversal = append(inorderTraversal, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return BSTIterator{inorderTraversal}

}


func (this *BSTIterator) Next() int {
    next := this.inorderTraversal[0]
    this.inorderTraversal = this.inorderTraversal[1:]
    return next
}


func (this *BSTIterator) HasNext() bool {
    return len(this.inorderTraversal) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */