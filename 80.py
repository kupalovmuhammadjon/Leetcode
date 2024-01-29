class Solution:
    def counter(self, nums,  n):
        return nums.count(n)
    def removeDuplicates(self, nums: List[int]) -> int:
        l = len(nums)
        i = 0
        res = []
        while i < len(nums):
            if self.counter(nums, nums[i]) > 2 and i != "_":
                nums.pop(i)
                l -= 1
            else:
                i += 1
        
        return l
        
