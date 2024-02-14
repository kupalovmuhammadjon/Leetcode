class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not len(nums):
            return 0
        nums = list(set(nums))
        nums.sort()

        mx = 0
        c = 1
        for i in range(1, len(nums)):
            if nums[i-1]+1 == nums[i]:
                c += 1
            else:
                mx = max(mx, c)
                c = 1
        mx = max(mx, c)
                    
        return mx