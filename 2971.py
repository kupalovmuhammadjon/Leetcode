class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        nums.sort()
        i = len(nums)-1
        while i > 1:
            if sum(nums[:i]) > nums[i]:
                return sum(nums[:i+1]) 
            i -= 1
        return -1