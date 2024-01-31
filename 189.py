class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        
        # for i in range(k):
        #     p = nums.pop()
        #     nums.insert(0, p)
        length = len(nums)
        k = k % length
        nums[:] = nums[-k:] + nums[:-k]
