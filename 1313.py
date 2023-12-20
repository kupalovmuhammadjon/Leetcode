class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        res = []
        i = 0

        while i < len(nums):
            freq, val = nums[i], nums[i + 1]

            for _ in range(freq):
                res.append(val)

            i += 2

        return res
