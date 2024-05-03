/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func levelOrder(root *Node) [][]int {
    answer := [][]int{}
    if root == nil{
        return answer
    }

    q := []*Node{root}

    for len(q) > 0{
        level := []int{}
        length := len(q)
        for i := 0; i < length; i++{
            curNode := q[0]
            level = append(level, curNode.Val)
            q = q[1:]
            q = append(q, curNode.Children...)
        }
        answer = append(answer, level)
    }
    return answer
}