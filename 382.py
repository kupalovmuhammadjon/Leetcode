# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:

    def __init__(self, head: Optional[ListNode]):
        self.head = head
        self.ls = []
        while self.head:
            self.ls.append(self.head)
            self.head = self.head.next

    def getRandom(self) -> int:
        choice = random.choice(self.ls)
        return choice.val
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(head)
# param_1 = obj.getRandom()