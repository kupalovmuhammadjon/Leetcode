# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        values = head
        ls = []
        while values:
            ls.append(values.val)
            values = values.next
        dummy = ListNode(0)
        dummy.next = head

        runner = dummy

        while runner and runner.next:
            if ls.count(runner.next.val) > 1:
                runner.next = runner.next.next
            else:
                runner = runner.next

        
        return dummy.next