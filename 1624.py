class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        
        mx = 0
        l = len(s)
        if len(set(s)) != l:
            for i in s:
                start = s.index(i)
                s = s[::-1]	
                stop = s.index(i)
                stop = l - stop - 2
                s = s[::-1]
                
                if mx < stop - start:
                    mx = stop - start
                    
            return mx
        else:
            return -1