/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func numComponents(head *ListNode, nums []int) int {
    count := 0
    mp := make(map[int]bool)
    for _, v := range nums{
        mp[v] = true
    }
    streak := 0
    for head != nil{
        if mp[head.Val]{
            streak++
        }else if streak > 0{
            count++
            streak = 0
        }
        head = head.Next
    }
    if streak > 0{
            count++
            streak = 0
        }
    return count
}