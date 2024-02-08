import heapq
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        sum = 0
        heapq.heapify(nums)
        ls = heapq.nlargest(2, nums)
        
        return (ls[0]-1) * (ls[1]-1)