class Solution:
    def sumZero(self, n: int) -> List[int]:
        
        res = []
        for i in range(n-1):
            
            res.append(i+1)
        res.append(-sum(res))
        return res    

# Done