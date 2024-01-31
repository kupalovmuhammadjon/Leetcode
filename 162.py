class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0
        elif len(nums) == 2:
            if nums == sorted(nums):
                return 1
            else:
                return 0
        if nums == sorted(nums, reverse=True):
            return 0

        for i in range(1, len(nums), 1):
            if i != len(nums)-1 and nums[i-1] < nums[i] > nums[i+1]:
                return i
        
        return len(nums)-1
            