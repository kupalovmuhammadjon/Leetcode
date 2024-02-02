# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        
        lkd = head
        res = ""
        while lkd:
            res += str(lkd.val)
            lkd = lkd.next
        if res == res[::-1]:
            return 1
        else:
            return 0

