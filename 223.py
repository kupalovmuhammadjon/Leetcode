# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:

        ls = []
        for i in range(len(lists)):
            head = lists[i]
            while head:
                ls.append(head.val)
                head = head.next
            
        ls.sort()
        head = ListNode(0)
        h = head
        for i in range(len(ls)):
            head.next = ListNode(ls[i])
            head = head.next

        return h.next


