class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        nums.sort()
        sm = 0
        for i in range(1, len(nums), 2):
            sm += min(nums[i-1], nums[i])

        return sm
        