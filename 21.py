# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        ls = []
        while list1:
            ls.append(list1.val)
            list1 = list1.next

        while list2:
            ls.append(list2.val)
            list2 = list2.next
        
        ls.sort()
        print(ls)
        if not ls:
            return list1

        head = ListNode(ls[0])
        st = head
        for i in range(1, len(ls)):
            head.next = ListNode(ls[i])
            head = head.next

        return st

        