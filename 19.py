# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0)
        dummy.next = head

        runner = dummy
        walker = dummy

        for _ in range(n+1):
            runner = runner.next

        while runner:
            runner = runner.next
            walker = walker.next
        
        walker.next = walker.next.next

        return dummy.next
        

        