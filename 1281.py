class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        ls = list(n)
        s = 1
        
        for i in ls:
            s *= i
            
        return s-sum(ls)
    
# Done