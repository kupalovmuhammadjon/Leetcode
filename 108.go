package main

import "fmt"

// TreeNode definition
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
	
// Helper function to build the BST
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return buildBST(nums, 0, len(nums)-1)
}

// Recursive function to build the BST
func buildBST(nums []int, left int, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left + (right-left)/2
    root := &TreeNode{Val: nums[mid]}
    root.Left = buildBST(nums, left, mid-1)
    root.Right = buildBST(nums, mid+1, right)
    return root
}

// Function to print the tree in order (for testing purposes)
func printInOrder(root *TreeNode) {
    if root != nil {
        printInOrder(root.Left)
        fmt.Printf("%d ", root.Val)
        printInOrder(root.Right)
    }
}

func main() {
    nums := []int{-10, -3, 0, 5, 9}
    root := sortedArrayToBST(nums)
    printInOrder(root)
}
