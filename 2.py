# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        n1 = ""
        n2 = ""

        while l1:
            n1 += str(l1.val)
            l1 = l1.next
        
        while l2:
            n2 += str(l2.val)
            l2 = l2.next
        n1 = n1[::-1]
        n2= n2[::-1]
        sm = eval(f"{n1} + {n2}")
        ls = []
        while sm:
            ls.append(sm%10)
            sm //= 10
        try:
            head = ListNode(ls[0])
            runner = head
        except:
            head = ListNode()
            return head
        for i in range(1, len(ls)):
            tp = ListNode(ls[i])
            runner.next = tp
            runner = runner.next

        return head