/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findFrequentTreeSum(root *TreeNode) []int {
    sumCount := map[int]int{}
    maxFreq := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil{
            return 0
        }

        leftSum := dfs(node.Left)
        rightSum := dfs(node.Right)
        sum := rightSum+leftSum+node.Val
        sumCount[sum]++
        if maxFreq < sumCount[sum]{
            maxFreq = sumCount[sum]
        }

        return sum
    }
    dfs(root)

    freqSums := []int{}
    for sum, freq := range sumCount{
        if freq == maxFreq{
            freqSums = append(freqSums, sum)
        }
    }
    return freqSums
}