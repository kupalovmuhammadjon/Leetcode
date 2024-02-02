# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        a = []
        b = []
        while headA:            
            a.append(headA)
            headA = headA.next

        while headB:
            b.append(headB)
            headB = headB.next

        for i in a:
            if i in b:
                return i
        
        return None
    