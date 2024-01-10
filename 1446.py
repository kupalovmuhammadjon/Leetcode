class Solution:
    def maxPower(self, s: str) -> int:
        mx = 1
        tempc = 1
        p = ""
        for i in s:
            if p == i:
                tempc += 1
            else:
                mx = max(tempc, mx)
                tempc = 1
                p = i
                
        mx = max(tempc, mx)
        
            
        return mx 