class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        
        mx1 = max(nums)
        nums.remove(mx1)
        mx2 = max(nums)
        nums.remove(mx2)
        mn1 = min(nums)
        nums.remove(mn1)
        mn2 = min(nums)
        
        return mx1*mx2 - mn1*mn2
    
# Done