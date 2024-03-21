/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := []int{}
	leaf2 := []int{}
	stack := []*TreeNode{root1}

	for len(stack) > 0 {
		left := stack[len(stack)-1].Left
		right := stack[len(stack)-1].Right
		if left == nil && right == nil {
			leaf1 = append(leaf1, stack[len(stack)-1].Val)
		}
		stack = stack[:len(stack)-1]
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
	}

	stack = []*TreeNode{root2}

	for len(stack) > 0 {
		left := stack[len(stack)-1].Left
		right := stack[len(stack)-1].Right
		if left == nil && right == nil {
			leaf2 = append(leaf2, stack[len(stack)-1].Val)
		}
		stack = stack[:len(stack)-1]
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
	}
	fmt.Println(leaf1)
	fmt.Println(leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}