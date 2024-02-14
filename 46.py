class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = list(product(nums, repeat=len(nums)))
        i = 0
        while i < len(res):
            if len(set(res[i])) != len(res[i]):
                res.pop(i)
            else:
                i += 1

        return res
