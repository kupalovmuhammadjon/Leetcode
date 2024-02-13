class Solution:
    def optimalDivision(self, nums: List[int]) -> str:
        if len(nums) == 1:
            return str(nums[0])
        elif len(nums) == 2:
            return f"{nums[0]}/{nums[1]}"
        res = str(nums[0]) + "/("
        for i in range(1, len(nums)):
            res += str(nums[i])
            if i != len(nums) - 1:
                res += "/"
        res += ")"

        return res