class Solution:
    def freqAlphabets(self, s: str) -> str:
        
        i = len(s)-1
        res = ""
        
        while i > 0:
            
            if s[i] == "#":
                n = int(s[i-2])*10 + int(s[i-1])
                res += chr(n + 96)
                i -= 3
            else:
                res += chr(int(s[i]) + 96)
                i -= 1
        return res[::-1]
            
# Done