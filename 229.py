class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        st = set(nums)
        res = []
        dif = len(nums) // 3
        for i in st:
            c = nums.count(i)
            if dif < c:
                res.append(i)
       
        return res