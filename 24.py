# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        dummy.next = head
        current = dummy

        while current.next and current.next.next:
            first_node = current.next
            second_node = current.next.next

            current.next = second_node
            first_node.next = second_node.next
            second_node.next = first_node

            current = current.next.next

        return dummy.next
