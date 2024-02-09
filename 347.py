class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        n = set(nums)
        dc = {}
        for i in n:
            dc[i] = nums.count(i)

        sorted_dc = sorted(dc.items(), key=lambda x: x[1], reverse=True)
        res = []
        for i in range(k):
            res.append(sorted_dc[i][0])
            
        return res
