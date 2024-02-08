import heapq
class Solution:
    def numberGame(self, nums: List[int]) -> List[int]:
        res = []
        heapq.heapify(nums)
        while nums:
            bob = heapq.heappop(nums)
            alice = heapq.heappop(nums)
            res.append(alice)
            res.append(bob)
        
        return res