class Solution:
    def rearrangeArray(self, nums: List[int]) -> List[int]:
       
        p = []
        n = []
        for i in nums:
            if i > 0:
               p.append(i)
            else:
                n.append(i)
        
        res = []
        for i in range(len(p)):
            res.append(p[i])
            res.append(n[i])
        
        return res
        