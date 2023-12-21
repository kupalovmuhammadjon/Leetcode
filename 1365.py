class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        
        res = []
        for i in nums:
            c = 0
            for j in range(len(nums)):
                if nums[j] < i:
                    c += 1
                    
            res.append(c)
        return res

# Done