class Solution:
    def generateTheString(self, n: int) -> str:
        
        res = ""
        
       
        if n % 2 == 0:
            res += (n - 1) * "a"
            res += "b"
        else:
            res += n * "a"
            
        return res
            
        