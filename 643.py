from typing import List

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        sm = sum(nums[:k])
        s = sm
        
        for i in range(k, len(nums)):
            sm += nums[i] - nums[i - k]
            s = max(s, sm)
        
        return s / k

# Done