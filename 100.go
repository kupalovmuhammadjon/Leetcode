/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        if q == nil{
            return true
        }
        return false
    }
    if q == nil {
        if p == nil{
            return true
        }
        return false
    }
    queue1 := []*TreeNode{p}
    queue2 := []*TreeNode{q}

    for len(queue1) > 0 && len(queue2) > 0{
        length := len(queue1)

        for i := 0; i < length; i++{
            cur1 := queue1[0]
            cur2 := queue2[0]
            queue1 = queue1[1:]
            queue2 = queue2[1:]
            if cur1.Val != cur2.Val{
                return false
            }
            if cur1.Left != nil{
                if cur2.Left != nil && cur1.Val == cur2.Val && cur1.Left.Val == cur2.Left.Val{
                    queue1 = append(queue1, cur1.Left)
                    queue2 = append(queue2, cur2.Left)
                }else{
                    return false
                }
            }else if cur2.Left != nil{
                return false
            }
            if cur1.Right != nil{
                if cur2.Right != nil && cur1.Val == cur2.Val && cur1.Right.Val == cur2.Right.Val{
                    queue1 = append(queue1, cur1.Right)
                    queue2 = append(queue2, cur2.Right)
                }else{
                    return false
                }
            }else if cur2.Right != nil{
                return false
            }

        } 
    }

    return true
}